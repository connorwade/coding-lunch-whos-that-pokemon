import fetchGraphQL from "./fetchGraphql"
import { POKEMON_DETAILS } from "./queries"

const getPokemonDetails = async (id) => {
    const query = POKEMON_DETAILS
    const res = await fetchGraphQL(
        query,
        {id: id},
        "pokemon_details",
    )
    const nodes = res.data.pokemon_v2_pokemonspecies[0].pokemon.nodes[0]

    const pokemonDetails = {
        name: nodes.name,
        flavorText: res.data.pokemon_v2_pokemonspecies[0].flavorText[0].flavor_text,
        weight: nodes.weight,
        height: nodes.height,
        stats: {
            hp: nodes.stats[0].base_stat,
            atk: nodes.stats[1].base_stat,
            def: nodes.stats[2].base_stat,
            spAtk: nodes.stats[3].base_stat,
            spDef: nodes.stats[4].base_stat,
            spd: nodes.stats[5].base_stat,
        },
        types: {
            type1: nodes.types[0].pokemon_v2_type.name,
            type2: nodes.types[1]?.pokemon_v2_type.name || null,
        }
    }

    return pokemonDetails
}

export default getPokemonDetails