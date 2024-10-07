use clap::Parser;

/// 命令行参数结构体
#[derive(Parser, Debug)]
#[command(name = "Swagger to Composable")]
#[command(about = "Generate TypeScript composable functions from swagger.json", long_about = None)]
pub struct Args {
    /// 输入 swagger.json 文件路径
    #[arg(short, long)]
    pub input: String,

    /// 输出 TypeScript 文件路径
    #[arg(short, long)]
    pub output: String,

    /// 输出的app名称
    #[arg(short, long)]
    pub app_name: String,

    /// 对端的app名称
    #[arg(short, long)]
    pub peer_name: String,
}
