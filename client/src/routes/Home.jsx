import { useState, useEffect } from 'react';
import "../styles/Home.css"
import { Container, Box, List, ListItem, ListItemText, Link } from '@mui/material';

function Home() {
  const [shops, setShops] = useState([])
  useEffect(() => {
    const fetchShops = async () => {
      try {
        const response = await fetch('http://localhost:9192/home');
        if (!response.ok) {
          throw new Error('Failed to fetch shops');
        }
        const data = await response.json();
        console.log(data)
        setShops(data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchShops();
  }, []);

  return (
    <Container>
      <Box>
        <h1>HairPlanet</h1>
        <List>
          {shops.map(shop => (<ListItem key={shop.id}><Link href={`/shop/${shop.name}`}><ListItemText primary={shop.name} secondary={shop.address + shop.city + shop.zip} /></Link></ListItem>))}
        </List>
      </Box>
    </Container>
  );
}

export default Home;
