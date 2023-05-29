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

window.updateBoard = function (boardState) {
    board = boardState;
    for (let i = 0; i < 9; i++) {
        switch (board[i]) {
            case 0:
                boardElement.children[i].textContent = "";
                break;
            case 1:
                boardElement.children[i].textContent = "O";
                boardElement.children[i].style.color = "saddlebrown";
                break;
            case 2:
                boardElement.children[i].textContent = "X";
                boardElement.children[i].style.color = "sandybrown";
                break;
        }
    }
};
