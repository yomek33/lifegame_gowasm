const numRows = 30;
const numCols = 30;

const gridContainer = document.getElementById("grid-container");
let coloredCells = [];
// グリッドを生成
function createGrid() {
  for (let i = 0; i < numRows; i++) {
    for (let j = 0; j < numCols; j++) {
      const cell = document.createElement("div");
      cell.className = "cell";
      cell.setAttribute("data-row", i);
      cell.setAttribute("data-col", j);
      gridContainer.appendChild(cell);
    }
  }
}

function cellClickHandler(event) {
  const targetCell = event.target;
  const row = parseInt(targetCell.getAttribute("data-row"));
  const col = parseInt(targetCell.getAttribute("data-col"));
  // console.log(`クリックされたセル - 行: ${row}, 列: ${col}`);
  targetCell.classList.add("clicked");

  coloredCells.push([row, col]);
}

// ボタンがクリックされたときの処理

document.getElementById("runButton").addEventListener("click", function () {
  console.log("Sending data to the server:", coloredCells);
});
// グリッドを生成
createGrid();

// セルのクリックイベントを追加
gridContainer.addEventListener("click", cellClickHandler);

//指定した座標のセルを塗る
function updateGrid(cells) {
  clearAllCellColors();
  if (cells && cells.length > 0) {
    const innerArray = cells[0];
    innerArray.forEach((cell) => {
      const row = cell[0];
      const col = cell[1];
      const cellElem = document.querySelector(
        `.cell[data-row="${row}"][data-col="${col}"]`
      );
      if (cellElem) {
        cellElem.classList.add("clicked");
      } else {
        console.log(`セルが見つかりません - 行: ${row}, 列: ${col}`);
      }
    });
  }
}
function clearAllCellColors() {
  const cellElems = document.querySelectorAll(".cell.clicked");
  cellElems.forEach((cellElem) => {
    cellElem.classList.remove("clicked");
  });
}
