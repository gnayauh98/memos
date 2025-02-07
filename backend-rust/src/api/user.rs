use std::sync::Arc;

use axum::{extract::State, routing::put, Router};

use crate::AppState;

pub fn new() -> Router<Arc<AppState>> {
    Router::new().route("/", put(registry))
}

async fn registry(State(app_state): State<Arc<AppState>>) -> String {
    // 提取用户名与密码
    let result: Result<(i64,), sqlx::Error> = sqlx::query_as("select $1")
        .bind(86_i64)
        .fetch_one(&app_state.pool)
        .await;

    if let Ok(row) = result {
        return format!("{}", row.0);
    }

    String::from("LPL")
}
