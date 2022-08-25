import fetchGraphQL from "./fetchGraphql"
import { POKEMON_NAMES } from "./queries"

const getPokemonNames = async (id) => {
    const query = POKEMON_NAMES
    const res = await fetchGraphQL(
        query,
        {},
        "pokemon_names",
    )

    let names = []
    for (const pokemon of res.data.pokemon_v2_pokemon) {
        names = [...names, pokemon.name]
    }

    return names

}

export default getPokemonNames