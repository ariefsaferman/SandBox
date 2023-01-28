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

const TOTAL_ITEMS_DATA = 100000;
const TOTAL_ITEM_VIEWS_DATA = 500000;
const PARAM_PER_INSERT = 10000;

async function insertItemsData() {
  let dataCount = 0;
  console.time("insert items");

  while (dataCount < TOTAL_ITEMS_DATA) {
    baseQuery = "INSERT INTO items (title, content, long_url, slug, created_at) VALUES ";
    placeholders = [];
    data = [];

    const NUM_OF_COL = 5;

    for (let i = 1; i <= PARAM_PER_INSERT * NUM_OF_COL; i += NUM_OF_COL) {
      placeholders.push(`($${i}, $${i + 1}, $${i + 2}, $${i + 3}, $${i + 4})`);

      let slug = faker.random.alphaNumeric(6, {casing: "mixed"});

      data.push(
        faker.random.words(2),
        faker.commerce.productDescription(),
        `${process.env.APP_REDIRECT_URL}/${slug}`,
        slug,
        faker.date.between('2022-01-01T00:00:00.000Z', '2023-01-01T00:00:00.000Z') 
      );
      dataCount += 1;
    }

    baseQuery += placeholders.join(", ");

    baseQuery += " ON CONFLICT DO NOTHING";

    await pool.query(baseQuery, data);
    console.log("data count: ", dataCount);
  }
  console.timeEnd("insert items");
  
  console.log("done");
}

async function insertItemViewsData() {
  let dataCount = 0;
  console.time("insert item views");

  while (dataCount < TOTAL_ITEM_VIEWS_DATA) {
    baseQuery = "INSERT INTO item_views (item_id, created_at) VALUES ";
    placeholders = [];
    data = [];

    const NUM_OF_COL = 2;

    for (let i = 1; i <= PARAM_PER_INSERT * NUM_OF_COL; i += NUM_OF_COL) {
      placeholders.push(`($${i}, $${i + 1})`);

      let slug = faker.random.alphaNumeric(6, {casing: "mixed"});

      data.push(
        faker.datatype.bigInt({min: 1, max: 10000}),
        faker.date.between('2022-01-01T00:00:00.000Z', '2023-01-01T00:00:00.000Z') 
      );
      dataCount += 1;
    }

    baseQuery += placeholders.join(", ");

    await pool.query(baseQuery, data);
    console.log("data count: ", dataCount);
  }
  console.timeEnd("insert item views");
  
  console.log("done");
}

(async function() {
  await Promise.all([
    insertItemsData(),
    insertItemViewsData(),
  ])
  pool.end();
})();

