use axum::{extract::Extension, routing::get, routing::post, Router};
mod api;
use sqlx::mysql::MySqlPoolOptions;

#[tokio::main]
async fn main() {
    let pool = MySqlPoolOptions::new()
        .max_connections(5)
        .connect("mysql://root:123456@192.168.1.2:3306/common")
        .await
        .unwrap();

    // build our application with a single route
    let app = Router::new()
        .route("/img/:id", get(api::get))
        .route("/img", post(api::save))
        .layer(Extension(pool));

    // run it with hyper on localhost:3000
    axum::Server::bind(&"0.0.0.0:3000".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}
