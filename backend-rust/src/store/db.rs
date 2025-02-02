use sqlx::postgres::PgPoolOptions;


#[derive(Clone)]
pub struct DB {
    //pub dsn: String,
    //pub config: Config,
    pub pool: sqlx::postgres::PgPool,
}

#[derive(Clone)]
pub struct Config {
    pub host: String,
    pub port: String,
    pub user: String,
    pub pass: String,
    pub db: String,
}

pub async fn new(config: Config) -> Result<DB, sqlx::Error> {

    let dsn = format!(
        "postgres://{}:{}@{}:{}/{}?sslmode=disable",
        config.user,
        config.pass,
        config.host,
        config.port,
        config.db,
    );

    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&dsn).await?;

    Ok(DB { pool })
}
