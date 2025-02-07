use std::sync::Arc;

use axum::{routing::get, Router};

use crate::AppState;

pub fn new() -> Router<Arc<AppState>> {
    Router::new().route("/", get(ping))
}

async fn ping() -> &'static str {
    "user ping"
}
