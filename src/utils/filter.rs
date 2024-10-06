use crate::utils::convert::to_camel_case_upper;
use inflector::cases::camelcase::to_camel_case; // 使用 inflector 库来转换 camelCase
use inflector::cases::snakecase::to_snake_case; // 使用 inflector 库来转换 snake_case
use std::collections::HashMap;
use tera::{Result, Value};

pub fn camel_case_filter(value: &Value, _args: &HashMap<String, Value>) -> Result<Value> {
    if let Some(s) = value.as_str() {
        let camel_cased = to_camel_case(s);
        Ok(Value::String(camel_cased))
    } else {
        Err(tera::Error::msg("Expected a string as input"))
    }
}

pub fn camel_case_initial_upper_filter(
    value: &Value,
    _args: &HashMap<String, Value>,
) -> Result<Value> {
    if let Some(s) = value.as_str() {
        let camel_cased = to_camel_case_upper(s);
        Ok(Value::String(camel_cased))
    } else {
        Err(tera::Error::msg("Expected a string as input"))
    }
}

pub fn snake_case_filter(value: &Value, _args: &HashMap<String, Value>) -> Result<Value> {
    if let Some(s) = value.as_str() {
        let snake_cased = to_snake_case(s);
        Ok(Value::String(snake_cased))
    } else {
        Err(tera::Error::msg("Expected a string as input"))
    }
}
