const POKEMON_DETAILS = `query pokemon_details($id: Int) {
    pokemon_v2_pokemonspecies(where: { id: { _eq: $id } }) {
      flavorText: pokemon_v2_pokemonspeciesflavortexts(limit: 1) {
        flavor_text
      }
      pokemon: pokemon_v2_pokemons_aggregate(limit: 1) {
        nodes {
          height
          name
          weight
          stats: pokemon_v2_pokemonstats {
            base_stat
            stat: pokemon_v2_stat {
              name
            }
          }
          types: pokemon_v2_pokemontypes {
            slot
            pokemon_v2_type {
              name
            }
          }
        }
      }
    }
  }`;

const POKEMON_NAMES = `query pokemon_names {
    pokemon_v2_pokemon(limit: 251) {
      name
    }
  }
  `;

export { POKEMON_DETAILS, POKEMON_NAMES };
