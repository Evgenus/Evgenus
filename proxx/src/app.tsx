import React, {useState} from 'react';
import {useImmer} from "use-immer";
import {Cell, CellData} from './cell';
import './app.css';

interface BoardProps {
    boardData: CellData[][];
    isGameWon: boolean;
    onClick: (x: number, y: number) => void;
}

function Board(props: BoardProps) {
    const {boardData, isGameWon, onClick} = props;
    return (
        <table>
            <tbody>
            {boardData.map((row) => {
                return <tr>
                    {row.map((cell) => {
                        return (<td>
                            <Cell cell={cell} isGameWon={isGameWon} onClick={onClick}/>
                        </td>)
                    })}
                </tr>
            })}
            </tbody>
        </table>
    );
}

function App() {
    const [started, setStarted] = useState(false);
    const [ended, setEnded] = useState(false);
    const [won, setWon] = useState(false);
    const [size, setSize] = useState(8);
    const [blackHoles, setBlackHoles] = useState(10);

    const [boardData, updateBoardData] = useImmer<CellData[][]>([]);

    const handleStart = () => {
        updateBoardData((draft) => {
            draft.length = size;
            for (let i = 0; i < size; i++) {
                draft[i] = [];
                for (let j = 0; j < size; j++) {
                    draft[i][j] = {
                        x: i,
                        y: j,
                        isOpen: false,
                        adjacentBlackHoles: 0,
                        isBlackHole: false,
                    };
                }
            }
        });

        globalThis.createGame(size, blackHoles);
        setStarted(true);
    }

    const handleRestart = () => {
        setStarted(false);
        setEnded(false);
    }

    const handleClick = (x: number, y: number) => {
        const cells = globalThis.clickCell(x, y);
        const status = globalThis.gameStatus();
        setEnded(status.isEnded);
        setWon(status.isWon);
        updateBoardData((draft) => {
            cells.forEach((cell) => {
                draft[cell.x][cell.y] = cell;
            });
        })
    }

    return (
        <div className="container">
            <div className="row">
                <div className="col-12 text-center">
                    <h1>Proxx parody</h1>
                </div>
            </div>
            {!started && <div className="row d-flex justify-content-center">
                <div className="col-4">
                    <div className="mb-3 row">
                        <label htmlFor="inputSize" className="col-sm-4 col-form-label">Board size:</label>
                        <div className="col-sm-8">
                            <input type="number" className="form-control" id="inputSize" value={size}
                                   onChange={
                                       (e) => setSize(parseInt(e.target.value))
                                   }/>
                        </div>
                    </div>
                    <div className="mb-3 row">
                        <label htmlFor="inputHoles" className="col-sm-4 col-form-label">Black holes:</label>
                        <div className="col-sm-8">
                            <input type="number" className="form-control" id="inputHoles" value={blackHoles}
                                   onChange={
                                       (e) => setBlackHoles(parseInt(e.target.value))
                                   }/>
                        </div>
                    </div>
                    <div className="row">
                        <div className="col-sm-12">
                            <button type="submit" className="btn btn-primary" onClick={handleStart}>Start</button>
                        </div>
                    </div>
                </div>
            </div>}
            {started && <div className="row">
                <div className="col-12 d-flex justify-content-center">
                    <Board boardData={boardData} isGameWon={won} onClick={handleClick}/>
                </div>
            </div>}
            {ended && <div className="row">
                <div className="col-12 d-flex justify-content-center">
                    <div className="text-center">
                        {won ? <div className="text-success">You won!</div>
                            : <div className="text-danger">You lost!</div>}
                        <button type="button" className="btn btn-primary" onClick={handleRestart}>Restart</button>
                    </div>
                </div>
            </div>}
        </div>
    );
}

export default App;
