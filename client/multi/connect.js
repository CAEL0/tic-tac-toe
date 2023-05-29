var conn;

if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onclose = function (e) {};
    conn.onmessage = function (e) {
        var boardState = JSON.parse(e.data);

        if (!Array.isArray(boardState)) {
            console.log("The board is not an array:", boardState);
            return;
        }
        if (boardState.length !== 9) {
            console.log("The board is not of length 9:", boardState);
            return;
        }
        if (
            !boardState.every((n) => {
                return n === 0 || n === 1 || n === 2;
            })
        ) {
            console.log("The board contains inappropriate element:", boardState);
            return;
        }
        window.updateBoard(boardState);
    };
}
