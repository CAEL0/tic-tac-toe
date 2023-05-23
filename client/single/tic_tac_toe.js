// 0: empty
// 1: player1
// 2: player2
let board = [0, 0, 0, 0, 0, 0, 0, 0, 0];
let currentPlayer = 1;
let gameOver = false;
const winningCombos = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6],
];
const boardElement = document.querySelector(".board");
const resultElement = document.querySelector(".result");

for (let i = 0; i < 9; i++) {
    const cell = document.createElement("div");
    cell.classList.add("cell");
    cell.dataset.index = i;
    cell.addEventListener("click", handleCellClick);
    boardElement.appendChild(cell);
}

function handleCellClick(event) {
    if (gameOver) return;

    const clickedCell = event.target;
    const cellIndex = parseInt(clickedCell.dataset.index);

    if (board[cellIndex] !== 0) return;

    board[cellIndex] = currentPlayer;
    clickedCell.textContent = currentPlayer === 1 ? "O" : "X";
    clickedCell.style.color = currentPlayer === 1 ? "saddlebrown" : "sandybrown";

    if (checkWin(currentPlayer)) {
        gameOver = true;
        resultElement.textContent = "Player " + currentPlayer + " wins";
        return;
    }
    if (checkDraw()) {
        gameOver = true;
        resultElement.textContent = "Draw";
        return;
    }
    currentPlayer = 3 - currentPlayer;
}

function checkWin(player) {
    return winningCombos.some((combo) => {
        return combo.every((index) => board[index] === player);
    });
}

function checkDraw() {
    return board.every((cell) => cell !== 0);
}

const resetButton = document.querySelector(".reset");

resetButton.addEventListener("click", () => {
    for (let i = 0; i < 9; i++) {
        board[i] = 0;
        boardElement.children[i].textContent = "";
    }
    currentPlayer = 1;
    gameOver = false;
    resultElement.textContent = "&nbsp;";
});
