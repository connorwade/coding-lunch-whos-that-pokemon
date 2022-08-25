import React from "react";
import { getPokemonDetails, getPokemonNames } from "./myApi";
import generateRandomInt from "./utils";

function App() {
  const [isLoading, setIsLoading] = React.useState(true);
  const [pokemonNames, setPokemonNames] = React.useState([]);
  const [pokemonDetails, setPokemonDetails] = React.useState({
    types: [],
    stats: [],
  });
  const [guessValue, setGuessValue] = React.useState("");
  const [hasWon, setHasWon] = React.useState(false);

  const handleGuessChange = (e) => {
    setGuessValue(e.target.value);
    console.log(guessValue);
  };

  let searchedPokemon = pokemonNames.filter((name) =>
    name.includes(guessValue)
  );

  const handleGuessSubmit = (e) => {
    if (guessValue === pokemonDetails.name) {
      setHasWon(true);
    } else {
      setHasWon(false);
    }

    e.preventDefault();
  };

  React.useEffect(() => {
    async function fetchData() {
      const id = generateRandomInt();
      const details = await getPokemonDetails(id);
      const names = await getPokemonNames();
      setPokemonNames(names);
      setPokemonDetails(details);
    }

    fetchData();
    setIsLoading(false);
  }, []);

  return (
    <div>
      <header>
        <h1>Whos' That Pokemon?</h1>
        <hr />
      </header>

      {isLoading ? (
        <p>Loading</p>
      ) : (
        <main>
          <h5>{pokemonDetails.flavorText}</h5>
          <h5>
            Type: {pokemonDetails.types.type1}{" "}
            {pokemonDetails.types?.type2 ? (
              <span>/ {pokemonDetails.types.type2}</span>
            ) : (
              <span></span>
            )}
          </h5>
          <h5>Height:{pokemonDetails.height}</h5>
          <h5>Weight:{pokemonDetails.weight}</h5>
          <h5>HP:{pokemonDetails.stats.hp}</h5>
          <h5>ATK:{pokemonDetails.stats.atk}</h5>
          <h5>DEF:{pokemonDetails.stats.def}</h5>
          <h5>SpATK:{pokemonDetails.stats.spAtk}</h5>
          <h5>SpDEF:{pokemonDetails.stats.spDef}</h5>
          <h5>SPD:{pokemonDetails.stats.spd}</h5>
          <h3 className="result">
            {hasWon ? (
              <span>Good Job! Play again sometime</span>
            ) : (
              <span>Keep on Guessing!</span>
            )}
          </h3>
        </main>
      )}
      <hr />
      <form onSubmit={handleGuessSubmit}>
        <label htmlFor="guessInput">So Who's That Pokemon?</label>
        <input id="guessInput" type="text" onChange={handleGuessChange} />
        <button>Guess</button>
      </form>
      <hr />
      {isLoading ? <p>Loading ...</p> : <List list={searchedPokemon} />}
    </div>
  );
}

function List({ list }) {
  return (
    <ol>
      {list.map((item) => (
        <li key={item} className="pokemonName">
          {item}
        </li>
      ))}
    </ol>
  );
}

export default App;
