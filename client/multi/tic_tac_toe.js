let board = [0, 0, 0, 0, 0, 0, 0, 0, 0];
let boardElement;
let resultElement;
let resetButton;

window.addEventListener("DOMContentLoaded", (e) => {
    boardElement = document.querySelector(".board");
    resultElement = document.querySelector(".result");
    resetButton = document.querySelector(".reset");
    for (let i = 0; i < 9; i++) {
        const cell = document.createElement("div");
        cell.classList.add("cell");
        cell.dataset.index = i;
        cell.addEventListener("click", handleCellClick);
        boardElement.appendChild(cell);
    }
});
