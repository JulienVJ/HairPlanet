import { useState } from 'react';
import Container from '@mui/material/Container';
import { Box, Button, TextField } from '@mui/material';

const FormHairdresser = () => {
    const storedUserId = localStorage.getItem('userId');
    const [haidresser, setHairdresser] = useState({
        firstName: "",
        lastName: "",
        ShopID: storedUserId,
    })

    const handleChange = (e) => {
        const { name, value } = e.target;
        setHairdresser(prevState => ({
            ...prevState,
            [name]: value
        }));
    }

    const handleSubmit = async () => {
        try {
            const requestOptions = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(haidresser)
            };

            const response = await fetch('http://localhost:9192/createHairdresser', requestOptions);

            if (!response.ok) {
                throw new Error('Failed to create hairdresser');
            }

        } catch (error) {
            console.error('Error creating hairdresser:', error);
        }
    }


    return (
        <Container>
            <Box>
                <h1>Création coiffeur</h1>
                <Box sx={{ display: "flex", padding: 3 }}>
                    <TextField
                        id="outlined-controlled"
                        label="firstName"
                        name="firstName"
                        value={haidresser.firstName}
                        onChange={handleChange}
                    />
                    <TextField
                        id="outlined-controlled"
                        label="lastName"
                        name="lastName"
                        value={haidresser.lastName}
                        onChange={handleChange}
                    />
                </Box>
            </Box>
            <Button variant="contained" onClick={handleSubmit}>Enregistrer</Button>
        </Container>
    )
}

export default FormHairdresser;
