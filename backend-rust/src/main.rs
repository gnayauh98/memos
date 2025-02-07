use std::sync::Arc;

use clap::Parser;
use store::db;

mod api;
mod store;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(long)]
    host: String,
    #[arg(long)]
    port: String,
    #[arg(long)]
    db: String,
    #[arg(long)]
    user: String,
    #[arg(long)]
    pass: String,
}

#[derive(Clone)]
struct AppState {
    pool: sqlx::PgPool,
}

#[tokio::main]
async fn main() -> Result<(), sqlx::Error> {
    let args = Args::parse();

    let store_db = store::db::new(db::Config {
        host: args.host.clone(),
        port: args.port.clone(),
        user: args.user.clone(),
        pass: args.pass.clone(),
        db: args.db.clone(),
    })
    .await?;

    let app = api::new().with_state(Arc::new(AppState {
        pool: store_db.pool,
    }));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:5623").await.unwrap();
    axum::serve(listener, app).await.unwrap();

    Ok(())
}
