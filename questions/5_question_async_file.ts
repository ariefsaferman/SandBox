/**
 * When working on this question, provide the type guards for JSON.parse() as well.
 * References:
 *   - https://stackoverflow.com/questions/38688822/how-to-parse-json-string-in-typescript/62438143#62438143
 *   - https://blog.logrocket.com/how-to-use-type-guards-typescript/
 *   - https://www.typescriptlang.org/docs/handbook/2/narrowing.html
 */

import fs from 'fs';

interface childJson {
  id: number;
  name: string;
  parent_id: number;
}

interface parentsJson {
  id: number; 
  first_name: string; 
  last_name: string;
  gender: string; 
  children: childJson[];
}

fs.readFile('./parents.json', 'utf8', (error: (Error | null), data: string) => {
  if (error) return console.log(error);
  // [
  //   {
  //     id: 1,
  //     first_name: 'Acong',
  //     last_name: 'Suherman',
  //     gender: 'Male'
  //   },
  //   {
  //     id: 2,
  //     first_name: 'Djoko',
  //     last_name: 'Santoso',
  //     gender: 'Male'
  //   },
  //   {
  //     id: 3,
  //     first_name: 'Sitorus',
  //     last_name: 'Hermanto',
  //     gender: 'Male'
  //   },
  //   {
  //     id: 4,
  //     first_name: 'Maria',
  //     last_name: 'Herlina',
  //     gender: 'Female'
  //   }
  // ]
  const parents: parentsJson[] = JSON.parse(data);
  fs.readFile('./children.json', 'utf8', (error: (Error | null), data: string) => {
    if (error) return console.log(error);
    // [
    //   {
    //     id: 1,
    //     name: 'Bobi',
    //     parent_id: 1
    //   },
    //   {
    //     id: 2,
    //     name: 'Susi',
    //     parent_id: 1
    //   },
    //   {
    //     id: 3,
    //     name: 'Hermione',
    //     parent_id: 2
    //   },
    //   {
    //     id: 4,
    //     name: 'Richard',
    //     parent_id: 2
    //   },
    //   {
    //     id: 5,
    //     name: 'Renata',
    //     parent_id: 2
    //   },
    //   {
    //     id: 6,
    //     name: 'Regina',
    //     parent_id: 3
    //   },
    //   {
    //     id: 7,
    //     name: 'Yuli',
    //     parent_id: 3
    //   },
    //   {
    //     id: 8,
    //     name: 'June',
    //     parent_id: 4
    //   }
    // ]
    const children = JSON.parse(data);
    const parentsWithChildren = parents.map((parent) => {
      children.forEach((child: childJson) => {
        if (!parent.children) parent.children = [];
        if (child.parent_id === parent.id) parent.children.push(child);
      });
      // [
      //   {
      //     id: 1,
      //     first_name: 'Acong',
      //     last_name: 'Suherman',
      //     gender: 'Male',
      //     children: [
      //       {
      //         id: 1,
      //         name: 'Bobi',
      //         parent_id: 1
      //       },
      //       {
      //         id: 2,
      //         name: 'Susi',
      //         parent_id: 1
      //       }
      //     ]
      //   },
      //   {
      //     id: 2,
      //     first_name: 'Djoko',
      //     last_name: 'Santoso',
      //     gender: 'Male',
      //     children: [
      //       {
      //         id: 3,
      //         name: 'Hermione',
      //         parent_id: 2
      //       },
      //       {
      //         id: 4,
      //         name: 'Richard',
      //         parent_id: 2
      //       },
      //       {
      //         id: 5,
      //         name: 'Renata',
      //         parent_id: 2
      //       }
      //     ]
      //   },
      //   {
      //     id: 3,
      //     first_name: 'Sitorus',
      //     last_name: 'Hermanto',
      //     gender: 'Male',
      //     children: [
      //       {
      //         id: 6,
      //         name: 'Regina',
      //         parent_id: 3
      //       },
      //       {
      //         id: 7,
      //         name: 'Yuli',
      //         parent_id: 3
      //       }
      //     ]
      //   },
      //   {
      //     id: 4,
      //     first_name: 'Maria',
      //     last_name: 'Herlina',
      //     gender: 'Female',
      //     children: [
      //       {
      //         id: 8,
      //         name: 'June',
      //         parent_id: 4
      //       }
      //     ]
      //   }
      // ]
      return parent;
    });
    console.log(JSON.stringify(parentsWithChildren, null, 2));
  });
});
