module.exports = function () {
  return {
    users: [
      {
        email: 'acong@mail.com',
        password:
          '$2a$10$uqtHyJb1IMna3g5jO.cR0eMKYy2uIhbxV9IzUQDmeK8QXg1Yvb.rG',
        id: 1,
      },
      {
        email: 'djoko@mail.com',
        password:
          '$2a$10$KQ5/PAx5AtTPd1Cknu9SeeRH62b8lWFIz070Y4vyBhh.BXI5QF176',
        id: 2,
      },
      {
        email: 'sitorus@mail.com',
        password:
          '$2a$10$rzUmCzrL0lGUAuf9kBf8t.0YZz2DXpQx6ojS1tL4sBItqigglekRW',
        id: 3,
      },
    ],
    books: [
      {
        id: 1,
        title: 'Coders at Work',
        author: 'Peter Seibel',
        image_url:
          'https://upload.wikimedia.org/wikipedia/en/thumb/b/be/Coders_at_work.jpg/220px-Coders_at_work.jpg',
        synopsis:
          "This is a who's who in the programming world - a fascinating look at how some of the best in the world do their work. Patterned after the best selling Founders at Work, the book represents two years of interviews with some of the top programmers of our times.",
        price: 150000,
      },
      {
        id: 2,
        title: 'The Mythical Man-Month',
        author: 'Fred Brooks',
        image_url:
          'https://upload.wikimedia.org/wikipedia/en/f/fd/Mythical_man-month_%28book_cover%29.jpg',
        synopsis:
          'Few books on software project management have been as influential and timeless as The Mythical Man-Month. With a blend of software engineering facts and thought-provoking opinions, Fred Brooks offers insight for anyone managing complex projects. These essays draw from his experience as project manager for the IBM System/360 computer family and then for OS/360, its massive software system. Now, 20 years after the initial publication of his book, Brooks has revisited his original ideas and added new thoughts and advice, both for readers already familiar with his work and for readers discovering it for the first time.',
        price: 250000,
      },
      {
        id: 3,
        title: 'Clean Code',
        author: 'Robert Cecil',
        image_url:
          'https://images.tokopedia.net/img/cache/700/product-1/2017/8/24/14481426/14481426_ef7e5cea-3bef-4a99-9cc0-803871c0a0fe_510_680.jpg',
        synopsis:
          "Even bad code can function. But if code isn't clean, it can bring a development organization to its knees. Every year, countless hours and significant resources are lost because of poorly written code. But it doesn't have to be that way.",
        price: 275000,
      },
      {
        id: 4,
        title: 'The Pragmatic Programmer',
        author: 'Andy Hunt and Dave Thomas',
        image_url:
          'https://images-na.ssl-images-amazon.com/images/I/41HXiIojloL._SX396_BO1,204,203,200_.jpg',
        synopsis:
          'Ward Cunningham Straight from the programming trenches, The Pragmatic Programmer cuts through the increasing specialization and technicalities of modern software development to examine the core process--taking a requirement and producing working, maintainable code that delights its users. It covers topics ranging from personal responsibility and career development to architectural techniques for keeping your code flexible and easy to adapt and reuse.',
        price: 100000,
      },
      {
        id: 5,
        title: 'Code Complete',
        author: 'Steve McConnell',
        image_url:
          'https://upload.wikimedia.org/wikipedia/en/5/58/Code_Complete_1st_edition.jpg',
        synopsis:
          'Widely considered one of the best practical guides to programming, Steve McConnell’s original CODE COMPLETE has been helping developers write better software for more than a decade. Now this classic book has been fully updated and revised with leading-edge practices—and hundreds of new code samples—illustrating the art and science of software construction. Capturing the body of knowledge available from research, academia, and everyday commercial practice, McConnell synthesizes the most effective techniques and must-know principles into clear, pragmatic guidance.',
        price: 80000,
      },
      {
        id: 6,
        title: 'The Art of Computer Programming',
        author: 'Donald Knuth',
        image_url:
          'https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/2018/9780201853926.jpg',
        synopsis:
          'Countless readers have spoken about the profound personal influence of Knuth’s work. Scientists have marveled at the beauty and elegance of his analysis, while ordinary programmers have successfully applied his “cookbook” solutions to their day-to-day problems. All have admired Knuth for the breadth, clarity, accuracy, and good humor found in his books.',
        price: 300000,
      },
    ],
  };
};
