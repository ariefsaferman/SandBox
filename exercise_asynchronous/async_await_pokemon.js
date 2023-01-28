/* eslint-disable prefer-template */
/* eslint-disable indent */
/* eslint-disable arrow-parens */
/* eslint-disable no-console */
/* eslint-disable no-unused-vars */
const fetchPokemon = async (limit) => {
  try {
    const response = await fetch(`https://pokeapi.co/api/v2/pokemon?limit=${limit}`);
    const data = await response.json();
    console.log('success');
    data.results.forEach((element, i) => {
        const name = element.name.charAt(0).toUpperCase() + element.name.slice(1);
        const idx = i + 1;
        console.log(idx + '. ' + name + ` - ${element.url}`);
    });
  } catch (error) {
    console.log('error');
    console.error(error);
  }
};

fetchPokemon(5);
