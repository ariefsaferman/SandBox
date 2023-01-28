/* eslint-disable @typescript-eslint/no-unused-vars */
import React, {useState} from 'react';
import './App.css';


const seats = [
  {id: "A1"},
  {id: "A2"},
  {id: "A3"},
  {id: "A4"},
  {id: "A5"}, 
  {id: "A6"},
  {id: "A7"},
  {id: "A8"},
  {id: "A9"},
  {id: "A10"},

  {id: "B1"},
  {id: "B2"},
  {id: "B3"},
  {id: "B4"},
  {id: "B5"}, 
  {id: "B6"},
  {id: "B7"},
  {id: "B8"},
  {id: "B9"},
  {id: "B10"},
];

const price: number = 50000;

export default function Seat() {
  const [selectedSeat, setSelectedSeat] = useState<string[]>([]); 
  const [unavailableSeat, setUnavailableSeat] = useState<string[]>([]); 

  function handleSeatClick(seatId: string) {
    if (selectedSeat.includes(seatId)) {
      setSelectedSeat(selectedSeat.filter((id) => id !== seatId)); 
    } else {
      setSelectedSeat([...selectedSeat, seatId]); 
    }
  }

  function handleBook(seatId: string[]) {
    setUnavailableSeat([...unavailableSeat, ...seatId]); 
    setSelectedSeat([]);
  }

  return (
    <div className='grid'>
      <h1 className='title__cinema'>Booking Cinema</h1>
      <div className="seat__grid">
        {
          seats.map((seat) => (
            <button 
            key={seat.id} 
            className={`seat ${selectedSeat.includes(seat.id) ? "selected" : "available"}`}
            onClick={() => handleSeatClick(seat.id)}
            disabled={unavailableSeat.includes(seat.id)}>
              {seat.id}
            </button>
          ))
        }
      </div>
      <div className='price'>Selected seats: {selectedSeat.length}, total = Rp {selectedSeat.length * price}</div>

      <div className="wrapper">
      <button className='book__btn' onClick={() => handleBook(selectedSeat)}>Book</button>
      </div>
      
    </div>
  )
}