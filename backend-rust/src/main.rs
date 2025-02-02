use std::sync::Arc;

use axum::{
    extract::State, routing::get, Router
};
use clap::Parser;
use store::db;

mod store;
mod api;

#[derive(Parser,Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(long)]
    host: String,
    #[arg(long)]
    port: String,
    #[arg(long)]
    db:String,
    #[arg(long)]
    user: String,
    #[arg(long)]
    pass:String,
}

#[derive(Clone)]
struct AppState {
    pool: sqlx::PgPool,
}

#[tokio::main]
async fn main() -> Result<(), sqlx::Error> {

    let args = Args::parse();

    let store_db = store::db::new(db::Config{
        host: args.host.clone(),
        port: args.port.clone(),
        user: args.user.clone(),
        pass: args.pass.clone(),
        db: args.db.clone(),
    }).await?;

    let app_state = Arc::new(AppState {
        pool: store_db.pool,
    });

    let auth_api = api::auth::new();

    let app = Router::new()
        .route("/", get(root))
        .nest("/user", auth_api)
        .with_state(app_state);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();

    Ok(())
}

async fn root(State(app):State<Arc<AppState>>) -> &'static str{
    let result :Result<(i64,), sqlx::Error> = sqlx::query_as("select $1")
        .bind(150_i64)
        .fetch_one(&app.pool).await;
    if let Ok(row) = result {
        println!("{}", row.0);
    }
    "hello, anqzi!"
}
