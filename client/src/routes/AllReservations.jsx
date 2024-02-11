import { useState, useEffect } from 'react';
import "../styles/AllReservations.css"

function Reservations() {
  const [reservations, setReservations] = useState([]);

  useEffect(() => {
    // Fetch data from the Golang API
    fetch('http://localhost:9192/reservations')
      .then(response => response.json())
      .then(data => {
        // Assuming the API returns an array of reservations
        setReservations(data);
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []); // Empty dependency array means this effect runs once after the first render

  console.log(reservations)
  return (
    <div>
      <h1 className='reservations-title'>Liste des réservations</h1>
      <div className='reservations-cards'>
              {reservations.map((n) => (
        <>
        <div className='reservation-card'>
          <p key={n.key}>Date de réservation : {n.date}</p>
          <p key={n.key}>Heure de réservation : {n.hours}</p>
          <p key={n.key}>Id du Salon : {n.shopId}</p>
          <p key={n.key}>Id de l'Employé : {n.employeeId}</p>
          <p key={n.key}>Disponible : {n.isReserved ? <span>Oui</span> : <span>Non</span>}</p>
        </div>
        </>
      ))}
      </div>
    </div>
  );
}

export default Reservations;
