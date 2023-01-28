/* eslint-disable no-shadow */
/* eslint-disable no-console */
/* eslint-disable no-undef */
/* eslint-disable no-template-curly-in-string */
/* eslint-disable no-unused-vars */
const wakeUp = (pName, next) => {
  console.log(`1. ${pName} Wake up at 5am in the morning`);
  return next(pName);
};

const prepare = (pName, next) => {
  console.log(`2. ${pName} prepare the tools to go to the office`);
  if (pName === 'Arief') {
    return next(pName);
  }
  return next('');
};

const goToWork = (pName, next) => {
  console.log(`3. Right after preparetion is done, ${pName} go to work`);
  return next(pName);
};

const work = (pName, next) => {
  console.log(`4. Finally ${pName} work on time`);
  if (pName === 'Arief') {
    return next(`${pName} success go to work `);
  }
  return next('Error: not same person that go to work');
};

const startRoutine = (pName) => {
  wakeUp(pName, (pName) => {
    prepare(pName, (pName) => {
      goToWork(pName, (pName) => {
        work(pName, (msg) => {
          console.log(msg);
        });
      });
    });
  });
};

startRoutine('Arief');
