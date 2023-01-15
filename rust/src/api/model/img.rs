use sqlx::{Error, MySqlPool};

#[derive(sqlx::FromRow)]
pub struct Img {
    id: i32,
    url: Option<String>,
    base64: Vec<u8>,
}

impl Img {
    pub async fn get_img_by_id(pool: &MySqlPool, id: i32) -> Result<String, Error> {
        Ok({
            let img = sqlx::query_as!(Img, "select * from img where id = ?", id)
                .fetch_one(pool)
                .await?;
            std::str::from_utf8(&(img.base64)).unwrap().to_string()
        })
    }
}
