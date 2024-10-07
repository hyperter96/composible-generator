use convert_case::{Case, Casing};
use std::collections::HashMap;
pub fn to_camel_case_upper(s: &str) -> String {
    let camel = s.to_case(Case::Camel);
    // 确保首字母大写
    let mut c = camel.chars();
    match c.next() {
        None => String::new(),
        Some(f) => f.to_uppercase().collect::<String>() + c.as_str(),
    }
}

// 函数：从 operationId 中提取下划线后的部分
pub fn extract_method_name(operation_id: &str) -> String {
    // 查找第一个下划线的位置
    if let Some(pos) = operation_id.find('_') {
        // 提取下划线后的子串
        let method = &operation_id[pos + 1..];
        to_camel_case_upper(method)
    } else {
        // 如果没有下划线，返回整个 operationId
        to_camel_case_upper(operation_id)
    }
}

pub fn insert_if_absent(map: &mut HashMap<String, String>, key: String, value: String) {
    match map.entry(key) {
        std::collections::hash_map::Entry::Vacant(entry) => {
            entry.insert(value); // 如果 key 不存在，插入 key 和 value
        }
        std::collections::hash_map::Entry::Occupied(_) => {
            // 如果 key 已存在，则什么都不做
        }
    }
}

pub fn insert_underscore_after_substring(s: &str, sub: &str) -> String {
    // 检查主字符串中是否包含子字符串
    if let Some(index) = s.find(sub) {
        // 计算插入位置，即子字符串的末尾
        let insert_pos = index + sub.len();

        // 将字符串分为两部分，拼接下划线
        let (before, after) = s.split_at(insert_pos);
        format!("{}{}{}", before, "_", after)
    } else {
        // 如果子字符串不在主字符串中，返回原字符串
        s.to_string()
    }
}

pub fn remove_substring(s: &str, sub: &str) -> String {
    // 查找子字符串的起始索引
    if let Some(index) = s.find(sub) {
        // 计算子字符串的末尾位置
        let start_of_remaining = index + sub.len();

        // 截取剩余的部分
        let remaining = &s[start_of_remaining..];
        return remaining.to_string(); // 返回剩余部分
    }
    // 如果子字符串不在主字符串中，返回原字符串
    s.to_string()
}
