import React from "react";
import "./App.css";
import "./style.css";

import { useState } from "react";



const price = 50000;

const seats = [
  { id: "A1" },
  { id: "A2" },
  { id: "A3" },
  { id: "A4" },
  { id: "A5" },
  { id: "A6" },
  { id: "A7" },
  { id: "A8" },
  { id: "A9" },
  { id: "A10" },
  { id: "B1" },
  { id: "B2" },
  { id: "B3" },
  { id: "B4" },
  { id: "B5" },
  { id: "B6" },
  { id: "B7" },
  { id: "B8" },
  { id: "B9" },
  { id: "B10" },
  { id: "C1" },
  { id: "C2" },
  { id: "C3" },
  { id: "C4" },
  { id: "C5" },
  { id: "C6" },
  { id: "C7" },
  { id: "C8" },
  { id: "C9" },
  { id: "C10" },
  { id: "D1" },
  { id: "D2" },
  { id: "D3" },
  { id: "D4" },
  { id: "D5" },
  { id: "D6" },
  { id: "D7" },
  { id: "D8" },
  { id: "D9" },
  { id: "D10" },
];

function App() {
return (
<div className="app">
<div className="screen">screen</div>

<div className="seats"></div>

<SeatSelection />
</div>
  );
}

function SeatSelection() {
const [selectedSeats, setSelectedSeats] = useState<string[]>([]);

const handleSeatClick = (seatId: string) => {
if (selectedSeats.includes(seatId)) {
setSelectedSeats(selectedSeats.filter((id) => id !== seatId));
    } else {
setSelectedSeats([...selectedSeats, seatId]);
    }
  };

return (
<div>
<div className="seat-grid">
{seats.map((seat) => (
<div
key={seat.id}
className={`seat ${
selectedSeats.includes(seat.id) ? "selected" : "available"
}`}
onClick={() => handleSeatClick(seat.id)}
>
{seat.id}
</div>
        ))}
</div>
<div className="price">Selected seats: {selectedSeats.length}, the total price would be Rp {selectedSeats.length*price}</div>
</div>
  );
}

export default App;