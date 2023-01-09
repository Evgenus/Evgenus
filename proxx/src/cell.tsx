import React, {useCallback} from "react";

interface CellData {
    x: number;
    y: number;
    isOpen: boolean;
    adjacentBlackHoles: number;
    isBlackHole: boolean;
}

interface CellProps {
    cell: CellData;
    isGameWon: boolean;
    onClick: (x: number, y: number) => void;
}

interface CellIconProps {
    cell: CellData;
    isGameWon: boolean;
}

function CellIcon({cell, isGameWon}: CellIconProps) {
    const {isOpen, isBlackHole, adjacentBlackHoles} = cell;
    if(!isOpen) {
        return <i className="bi bi-square-fill"/>;
    }
    if(isBlackHole) {
        const colorClass = isGameWon ? 'text-success' : 'text-danger';
        return <i className={`bi bi-x-square-fill ${colorClass}`}/>;
    }
    switch (adjacentBlackHoles) {
        case 0: return <i className="bi bi-square"/>;
        case 1: return <i className="bi bi-1-square"/>
        case 2: return <i className="bi bi-2-square"/>
        case 3: return <i className="bi bi-3-square"/>
        case 4: return <i className="bi bi-4-square"/>
        case 5: return <i className="bi bi-5-square"/>
        case 6: return <i className="bi bi-6-square"/>
        case 7: return <i className="bi bi-7-square"/>
        case 8: return <i className="bi bi-8-square"/>
    }
}

function Cell(props: CellProps) {
  const { cell, isGameWon, onClick } = props;
  const { x, y } = cell;

  const handleClick = useCallback(() => {
      if(!cell.isOpen) {
          onClick(x, y);
      }
  }, [x, y, onClick]);

  return (
    <div onClick={handleClick}>
        <CellIcon cell={cell} isGameWon={isGameWon}/>
    </div>
  );
}

export {
    Cell,
    CellData,
};