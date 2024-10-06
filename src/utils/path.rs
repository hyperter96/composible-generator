use std::fs;
use std::path::{Path, PathBuf};

/// 获取文件的绝对路径
fn get_absolute_path(file_name: &str) -> Option<PathBuf> {
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
fn get_file_prefix(file_name: &str) -> Option<String> {
    let path = Path::new(file_name);
    if let Some(stem) = path.file_stem() {
        return Some(stem.to_string_lossy().to_string());
    }
    None
}
