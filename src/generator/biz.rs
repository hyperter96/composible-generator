use crate::models::common::Args;
use crate::models::swagger::{
    ComposableFunction, Definition, FunctionParameter, FunctionReturn, OpenAPI, Schema,
};
use crate::utils::convert::{extract_method_name, insert_if_absent, to_camel_case_upper};
use crate::utils::filter::{camel_case_filter, camel_case_initial_upper_filter, snake_case_filter};
use inflector::cases::camelcase::to_camel_case;
use std::collections::HashMap;
use std::fs::File;
use std::io::Read;
use tera::{Context, Tera};

fn get_return_type(schema: &Schema) -> String {
    if let Some(ref_path) = &schema.ref_path {
        // 提取引用类型名
        if let Some(pos) = ref_path.find("v1") {
            to_camel_case_upper(&ref_path[pos + 2..])
        } else if let Some(pos) = ref_path.find("definitions/") {
            to_camel_case_upper(&ref_path[pos + 12..])
        } else {
            "any".to_string()
        }
    } else if let Some(format) = &schema.format {
        match format.as_str() {
            "int32" => "int32".to_string(),
            "int64" => "int64".to_string(),
            "float" => "float32".to_string(),
            "double" => "float64".to_string(),
            _ => "any".to_string(),
        }
    } else if let Some(ty) = &schema.ty {
        match ty.as_str() {
            "string" => "string".to_string(),
            "boolean" => "bool".to_string(),
            "array" => {
                if schema.schema.is_none() {
                    "*structpb.ListValue".to_string()
                } else {
                    "[]".to_string()
                        + &get_return_type(
                            schema
                                .schema
                                .as_ref()
                                .unwrap_or(&Box::new(Schema::default())),
                        )
                }
            }
            "object" => {
                if schema.schema.is_none() {
                    "*structpb.Struct".to_string()
                } else {
                    get_return_type(
                        schema
                            .schema
                            .as_ref()
                            .unwrap_or(&Box::new(Schema::default())),
                    )
                }
            }
            _ => "any".to_string(),
        }
    } else {
        "any".to_string()
    }
}

fn get_item_type(
    item_map: &mut HashMap<String, Vec<FunctionReturn>>,
    properties: &HashMap<String, Schema>,
    definitions: &HashMap<String, Definition>,
) {
    let mut record_map = HashMap::new();
    for (_property_name, property_schema) in properties.iter() {
        if let Some(ref_path) = &property_schema.ref_path {
            let return_ty = if let Some(pos) = ref_path.find("v1") {
                to_camel_case_upper(&ref_path[pos + 2..])
            } else if let Some(pos) = ref_path.find("definitions/") {
                to_camel_case_upper(&ref_path[pos + 12..])
            } else {
                "any".to_string()
            };
            let ty = if let Some(pos) = ref_path.find("definitions/") {
                ref_path[pos + 12..].to_string()
            } else {
                "any".to_string()
            };
            definitions
                .get(&ty)
                .unwrap()
                .properties
                .as_ref()
                .unwrap()
                .iter()
                .for_each(|(k, v)| {
                    item_map
                        .entry(return_ty.clone())
                        .or_default()
                        .push(FunctionReturn {
                            name: k.to_string(),
                            ty: get_return_type(v),
                        });
                });
            get_item_type(
                item_map,
                definitions.get(&ty).unwrap().properties.as_ref().unwrap(),
                definitions,
            );
        } else if let Some(schema) = &property_schema.schema {
            if let Some(ref_path) = &schema.ref_path {
                let return_ty = if let Some(pos) = ref_path.find("v1") {
                    to_camel_case_upper(&ref_path[pos + 2..])
                } else if let Some(pos) = ref_path.find("definitions/") {
                    to_camel_case_upper(&ref_path[pos + 12..])
                } else {
                    "any".to_string()
                };
                let ty = if let Some(ref_path) = &schema.ref_path {
                    if let Some(pos) = ref_path.find("definitions/") {
                        ref_path[pos + 12..].to_string()
                    } else {
                        continue;
                    }
                } else {
                    continue;
                };
                if record_map.contains_key(&return_ty) {
                    continue;
                }
                definitions
                    .get(&ty)
                    .unwrap()
                    .properties
                    .as_ref()
                    .unwrap()
                    .iter()
                    .for_each(|(k, v)| {
                        item_map
                            .entry(return_ty.clone())
                            .or_default()
                            .push(FunctionReturn {
                                name: k.to_string(),
                                ty: get_return_type(v),
                            });
                    });
                record_map.insert(return_ty.clone(), true);
            };
        } else {
            continue;
        }
    }
}

/// 生成业务代码
pub fn generate_biz_from_swagger(args: &Args) -> Result<(), Box<dyn std::error::Error>> {
    // Step 2: 读取 OpenAPI 文件
    let mut file = File::open(&args.input)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    let openapi: OpenAPI = serde_json::from_str(&contents)?;

    // Step 3: 初始化 Tera 并注册自定义过滤器
    let mut tera = Tera::new("templates/**/*")?;

    // 注册自定义的 camel_case 过滤器
    tera.register_filter("camel_case", camel_case_filter);
    tera.register_filter("camel_case_upper", camel_case_initial_upper_filter);
    // 注册自定义的 snake_case 过滤器
    tera.register_filter("snake_case", snake_case_filter);

    // Step 4: 遍历路径和操作，提取函数名和参数
    let mut composable_functions = Vec::new();
    let mut service_name = None;
    let mut found_struct = false;
    if let Some(tag) = openapi.tags.first() {
        if let Some(tag_name) = tag.get("name") {
            service_name = Some(tag_name.clone());
        }
    }
    for (path, path_item) in openapi.paths.iter() {
        for (method, operation_opt) in &[
            ("get", &path_item.get),
            ("post", &path_item.post),
            ("put", &path_item.put),
            ("delete", &path_item.delete),
        ] {
            if let Some(operation) = operation_opt {
                // 获取函数名(没有Service的字符串)
                let func_name = match &operation.operationId {
                    Some(id) => extract_method_name(id),
                    None => {
                        // 备选方案：方法名 + 路径
                        let sanitized_path =
                            path.replace("/", "_").replace("{", "").replace("}", "");
                        format!("{}_{}", method, sanitized_path)
                    }
                };

                // 获取描述
                let description = operation.summary.clone().unwrap_or_else(|| "".to_string());

                // 处理参数
                let mut req_params = Vec::new();
                let res_params = Vec::new();
                let mut request_struct = None;
                if let Some(parameters) = &operation.parameters {
                    for param in parameters {
                        let ty = match &param.ty {
                            Some(ty) => match ty.as_str() {
                                "string" => "string".to_string(),
                                "integer" => {
                                    if let Some(format) = &param.format {
                                        match format.as_str() {
                                            "int32" => "int32".to_string(),
                                            "int64" => "int64".to_string(),
                                            _ => "any".to_string(),
                                        }
                                    } else {
                                        "int32".to_string()
                                    }
                                }
                                "boolean" => "bool".to_string(),
                                "array" => {
                                    "[]".to_string()
                                        + &get_return_type(param.schema.as_ref().unwrap())
                                }
                                "object" => {
                                    if param.schema.is_none() {
                                        "*structpb.Struct".to_string()
                                    } else {
                                        get_return_type(param.schema.as_ref().unwrap())
                                    }
                                }
                                _ => "any".to_string(),
                            },
                            None => {
                                if let Some(schema) = &param.schema {
                                    get_return_type(schema)
                                } else {
                                    "any".to_string()
                                }
                            }
                        };
                        req_params.push(FunctionParameter {
                            name: to_camel_case(&param.name),
                            ty,
                            required: param.required.unwrap_or(false),
                            location: param.location.clone(),
                        });
                    }

                    request_struct = Some(format!(
                        "{}Request",
                        extract_method_name(
                            &operation
                                .operationId
                                .clone()
                                .unwrap_or_else(|| func_name.clone())
                        )
                    ));
                }

                // 获取返回类型（简化处理，只考虑 200 响应）
                let return_type = if let Some(responses) = &operation.responses {
                    if let Some(response) = responses.get("200") {
                        if let Some(content) = &response.content {
                            if let Some(media_type) = content.get("application/json") {
                                if let Some(schema) = &media_type.schema {
                                    Some(get_return_type(schema))
                                } else {
                                    Some("any".to_string())
                                }
                            } else {
                                Some("any".to_string())
                            }
                        } else if let Some(schema) = &response.schema {
                            Some(get_return_type(schema))
                        } else {
                            Some("any".to_string())
                        }
                    } else {
                        Some("any".to_string())
                    }
                } else {
                    Some("any".to_string())
                };

                composable_functions.push(ComposableFunction {
                    name: func_name,
                    request_struct,
                    description,
                    request_parameters: req_params,
                    response_parameters: res_params,
                    return_type,
                    reveal_struct: true,
                });
            }
        }
    }

    // Step 5: 遍历definitions，提取结构体名
    let mut item_map = HashMap::new();
    let mut resp_map = HashMap::new();
    for (return_ty, schema) in openapi.definitions.iter() {
        let return_type = to_camel_case_upper(return_ty.split("v1").last().unwrap_or("any"));
        if let Some(properties) = &schema.properties {
            composable_functions.iter_mut().for_each(|func| {
                if Some(return_type.clone()) == func.return_type
                    && !resp_map.contains_key(return_type.clone().as_str())
                {
                    func.response_parameters = properties
                        .iter()
                        .map(|(property_name, property_schema)| FunctionReturn {
                            name: to_camel_case_upper(property_name),
                            ty: get_return_type(property_schema),
                        })
                        .collect();
                    insert_if_absent(&mut resp_map, return_type.clone(), func.name.clone());
                    get_item_type(&mut item_map, properties, &openapi.definitions);
                } else if Some(return_type.clone()) == func.return_type
                    && resp_map.contains_key(return_type.clone().as_str())
                {
                    func.reveal_struct = false
                }
            });
        }
    }

    // for_each闭包无法捕获外部可变变量
    for func in composable_functions.iter_mut() {
        for param in func.response_parameters.iter_mut() {
            if &param.ty == "*structpb.Struct" || &param.ty == "*structpb.ListValue" {
                found_struct = true;
                break;
            }
        }
        if found_struct {
            break;
        }
    }

    for (_key, item) in item_map.iter_mut() {
        for param in item.iter_mut() {
            if &param.ty == "*structpb.Struct" || &param.ty == "*structpb.ListValue" {
                found_struct = true;
                break;
            }
        }
        if found_struct {
            break;
        }
    }

    // Step 6: 准备模板上下文并渲染
    let mut context = Context::new();
    context.insert("functions", &composable_functions);
    context.insert("service", &service_name);
    context.insert("items", &item_map);
    context.insert("found_struct", &found_struct);

    // Step 7: 渲染模板
    let rendered = tera.render("biz_template.tera", &context)?;

    // Step 8: 写入生成的 biz_template 文件
    let mut output_file = File::create(&args.output)?;
    use std::io::Write;
    output_file.write_all(rendered.as_bytes())?;

    println!("biz层函数已生成到 output/composables/biz.go");

    Ok(())
}
