import { useState, useEffect } from 'react';
import "../styles/Home.css"

function Home() {
  const [shops, setShops] = useState([])
  useEffect(() => {
    const fetchShops = async () => {
      try {
        const response = await fetch('http://localhost:9192/getShops');
        if (!response.ok) {
          throw new Error('Failed to fetch shops');
        }
        const data = await response.json();
        setShops(data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchShops();
  }, []);

  return (
    <div>
      <h1>HairPlanet</h1>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
      {shops.map(shop => (<ul key={shop.id}>{shop.name}</ul>))}
    </div>
  );
}

export default Home;
