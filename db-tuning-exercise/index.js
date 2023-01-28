const faker = require("@faker-js/faker").faker;
const Pool = require("pg").Pool;
require('dotenv').config()

// setting database
const pool = new Pool({
  user: process.env.DB_USER,
  host: process.env.DB_HOST,
  database: process.env.DB_NAME,
  password: process.env.DB_PASSWORD,
  port: process.env.DB_PORT,
});

const TOTAL_CATEGORIES_DATA = 10000;
const TOTAL_TAGS_DATA = 10000;
const TOTAL_ARTICLES_DATA = 100000;
const TOTAL_ARTICLE_TAGS_DATA = 50000;
const PARAM_PER_INSERT = 10000;

async function insertTagsData() {
  let dataCount = 0;
  console.time("insert tags");

  while (dataCount < TOTAL_TAGS_DATA) {
    baseQuery = "INSERT INTO tags (name, description) VALUES ";
    placeholders = [];
    data = [];

    const NUM_OF_COL = 2;

    for (let i = 1; i <= PARAM_PER_INSERT * NUM_OF_COL; i += NUM_OF_COL) {
      placeholders.push(`($${i}, $${i + 1})`);
      data.push(
        faker.random.words(2),
        faker.commerce.productDescription(),
      );
      dataCount += 1;
    }

    baseQuery += placeholders.join(", ");

    await pool.query(baseQuery, data);
    console.log("data count: ", dataCount);
  }
  console.timeEnd("insert tags");
  
  console.log("done");
}

async function insertCategoriesData() {
  let dataCount = 0;
  console.time("insert categories");

  while (dataCount < TOTAL_CATEGORIES_DATA) {
    baseQuery = "INSERT INTO categories (name, description) VALUES ";
    placeholders = [];
    data = [];

    const NUM_OF_COL = 2;

    for (let i = 1; i <= PARAM_PER_INSERT * NUM_OF_COL; i += NUM_OF_COL) {
      placeholders.push(`($${i}, $${i + 1})`);
      data.push(
        faker.random.words(3),
        faker.commerce.productDescription(),
      );
      dataCount += 1;
    }

    baseQuery += placeholders.join(", ");

    await pool.query(baseQuery, data);
    console.log("data count: ", dataCount);
  }
  console.timeEnd("insert categories");
  
  console.log("done");
}

async function insertArticlesData() {
  let dataCount = 0;
  console.time("insert articles");

  const status = [
    'PUBLISHED',
    'DRAFT',
    'ARCHIVED'
  ];

  while (dataCount < TOTAL_ARTICLES_DATA) {
    baseQuery = "INSERT INTO articles (title, content, status, category_id) VALUES ";
    placeholders = [];
    data = [];

    const NUM_OF_COL = 4;

    for (let i = 1; i <= PARAM_PER_INSERT * NUM_OF_COL; i += NUM_OF_COL) {
      placeholders.push(`($${i}, $${i + 1}, $${i + 2}, $${i + 3})`);
      data.push(
        faker.random.words(5),
        faker.commerce.productDescription(),
        status[Math.floor(Math.random() * 3)],
        Math.floor(Math.random() * TOTAL_CATEGORIES_DATA),
      );
      dataCount += 1;
    }

    baseQuery += placeholders.join(", ");

    await pool.query(baseQuery, data);
    console.log("data count: ", dataCount);
  }
  console.timeEnd("insert articles");
  console.log("done");
}

async function insertArticleTagsData() {
  let dataCount = 0;
  console.time("insert article tags");

  while (dataCount < TOTAL_ARTICLE_TAGS_DATA) {
    baseQuery = "INSERT INTO article_tags (article_id, tag_id) VALUES ";
    placeholders = [];
    data = [];

    const NUM_OF_COL = 2;

    for (let i = 1; i <= PARAM_PER_INSERT * NUM_OF_COL; i += NUM_OF_COL) {
      placeholders.push(`($${i}, $${i + 1})`);
      data.push(
        Math.floor(Math.random() * TOTAL_ARTICLES_DATA),
        Math.floor(Math.random() * TOTAL_TAGS_DATA),
      );
      dataCount += 1;
    }

    baseQuery += placeholders.join(", ");

    await pool.query(baseQuery, data);
    console.log("data count: ", dataCount);
  }
  console.timeEnd("insert article tags");
  
  console.log("done");
}

(async function() {
  await Promise.all([
    insertTagsData(),
    insertCategoriesData(),
    insertArticlesData(),
    insertArticleTagsData(),
  ])
  pool.end();
})();

