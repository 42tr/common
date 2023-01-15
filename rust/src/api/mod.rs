use axum::extract::{Extension, Path};
use axum::http::{HeaderMap, HeaderValue};
use base64::decode;
use sqlx::MySqlPool;
mod error;
mod model;

pub async fn save() -> Vec<u8> {
    let a = "hello";
    a.to_string().into_bytes()
}

pub async fn get(
    Path(id ): Path<i32>,
    Extension(pool): Extension<MySqlPool>,
) -> Result<(HeaderMap, Vec<u8>), error::CustomError> {
    let mut headers = HeaderMap::new();
    headers.insert("content-type", HeaderValue::from_static("image/jpeg"));

    let code = model::img::Img::get_img_by_id(&pool, id).await.expect("get img error");
    let dec = decode(code).expect("decode error");

    Ok((headers, dec))
}
