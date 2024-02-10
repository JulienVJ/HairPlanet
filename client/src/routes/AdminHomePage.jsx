import { useState, useEffect } from 'react';
import "../styles/AdminHomePage.css"

function Users() {
  const [names, setNames] = useState([]);

  useEffect(() => {
    // Fetch data from the Golang API
    fetch('http://localhost:9192/users')
      .then(response => response.json())
      .then(data => {
        // Assuming the API returns an array of names
        setNames(data);
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []); // Empty dependency array means this effect runs once after the first render

  return (
    <div>
      <h1>HairPlanet</h1>
      {names.map((n) => (
        <>
          <p key={n.key}>{n.name}</p>
          <p key={n.key}>{n.email}</p>
        </>

      ))}
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </div>
  );
}

export default Users;
