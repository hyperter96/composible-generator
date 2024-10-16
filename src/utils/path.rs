use std::fs;
use std::path::{Path, PathBuf};

/// 获取文件的绝对路径
pub fn get_absolute_path(file_name: &str) -> Option<PathBuf> {
    let path = Path::new(file_name);
    if path.exists() {
        // 转换为绝对路径
        match path.canonicalize() {
            Ok(abs_path) => Some(abs_path),
            Err(_) => None,
        }
    } else {
        None
    }
}

/// 获取文件的前缀名称（不包含扩展名）
pub fn get_file_prefix(file_name: &str) -> Option<String> {
    let path = Path::new(file_name);
    if let Some(file_name) = path.file_name() {
        let file_name_str = file_name.to_str().unwrap_or("");
        if let Some(index) = file_name_str.find('.') {
            return Some(file_name_str[..index].to_string());
        }
    }
    None
}
