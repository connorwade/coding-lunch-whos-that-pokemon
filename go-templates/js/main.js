const pokemonNames = document.querySelectorAll(".pokemonName")
const guessInput = document.querySelector("#guessInput")
const form = document.querySelector('form')
const result = document.querySelector('.result')

guessInput.addEventListener('input', filterNames)

form.addEventListener('submit', fetchName)

function filterNames(e) {
    for (const name of pokemonNames) {
        if (!name.innerText.includes(e.target.value)) {
            name.style.display = "none"
        } else {
            name.style.display = "list-item"
        }
    }
}

function fetchName(e) {
    e.preventDefault();
    const data =
    {
        guess: guessInput.value
    }

    let jsonData = JSON.stringify(data)

    fetch("/answer", {
        method: "post",
        body: jsonData,
    }).then(resp => resp.json())
        .then(data => {
            if(data.winlose === 'lose') {
                wrongGuess();
            } else {
                correctGuess();
            }
        })

    return false
}

function wrongGuess() {
    result.innerText = "Wrong, try again"
}

function correctGuess() {
    result.innerText = "Correct! Play again sometime"
}