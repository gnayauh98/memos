use std::sync::Arc;

use axum::Router;

use crate::AppState;

pub mod auth;
pub mod user;

pub fn new() -> Router<Arc<AppState>> {
    Router::new()
        .nest("/api/auth", auth::new())
        .nest("/api/user", user::new())
}
