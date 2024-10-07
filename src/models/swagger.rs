use serde::{Deserialize, Serialize};
use std::collections::HashMap;

// 定义 OpenAPI 规范的必要部分
#[derive(Debug, Deserialize)]
pub struct OpenAPI {
    pub tags: Vec<HashMap<String, String>>,
    pub paths: HashMap<String, PathItem>,
    pub definitions: HashMap<String, Definition>,
}

// 定义Definition部分
#[derive(Debug, Deserialize)]
pub struct Definition {
    #[serde(rename = "type")]
    pub ty: Option<String>,
    #[serde(rename = "properties")]
    pub properties: Option<HashMap<String, Schema>>,
}

#[derive(Debug, Deserialize)]
pub struct PathItem {
    #[serde(default)]
    pub get: Option<Operation>,
    #[serde(default)]
    pub post: Option<Operation>,
    #[serde(default)]
    pub put: Option<Operation>,
    #[serde(default)]
    pub delete: Option<Operation>,
    // 可根据需要添加其他 HTTP 方法
}

#[derive(Debug, Deserialize)]
pub struct Operation {
    pub summary: Option<String>,
    pub operationId: Option<String>,
    pub parameters: Option<Vec<Parameter>>,
    pub responses: Option<HashMap<String, Response>>,
}

#[derive(Debug, Deserialize)]
pub struct Parameter {
    pub name: String,
    #[serde(rename = "in")]
    pub location: String,
    #[serde(rename = "required")]
    pub required: Option<bool>,
    #[serde(rename = "type")]
    pub ty: Option<String>,
    #[serde(rename = "format")]
    pub format: Option<String>,
    pub schema: Option<Schema>,
}

#[derive(Debug, Deserialize, Default)]
pub struct Schema {
    #[serde(rename = "type")]
    pub ty: Option<String>,
    #[serde(rename = "format")]
    pub format: Option<String>,
    #[serde(rename = "$ref")]
    pub ref_path: Option<String>,
    #[serde(rename = "items")]
    pub schema: Option<Box<Schema>>,
}

#[derive(Debug, Deserialize)]
pub struct Response {
    pub description: Option<String>,
    pub content: Option<HashMap<String, MediaType>>,
    pub schema: Option<Schema>,
}

#[derive(Debug, Deserialize)]
pub struct MediaType {
    pub schema: Option<Schema>,
}

// 定义函数的入参
#[derive(Debug, Serialize)]
pub struct FunctionParameter {
    pub name: String,
    pub ty: String,
    pub sub_ty: String,
    pub tp: String,
    pub sub_tp: String,
    pub required: bool,
    pub location: String,
}

// 定义函数的返回值
#[derive(Debug, Serialize)]
pub struct FunctionReturn {
    pub name: String,
    pub ty: String,
    pub sub_ty: String,
    pub sub_proto_ty: String,
    pub tp: String,
    pub sub_tp: String,
}

// 定义传递给模板的数据结构
#[derive(Debug, Serialize)]
pub struct ComposableFunction {
    pub name: String,
    pub request_struct: Option<String>, // 如果使用结构体参数
    pub description: String,
    pub request_parameters: Vec<FunctionParameter>,
    pub response_parameters: Vec<FunctionReturn>,
    pub reveal_struct: bool,
    pub return_type: Option<String>,
}
