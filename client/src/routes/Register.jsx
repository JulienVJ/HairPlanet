import React from 'react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Checkbox, Grid, Button, FormControlLabel, TextField, Typography } from '@mui/material';

const Register = () => {
    const navigateTo = useNavigate();
    const [registrationData, setRegistrationData] = useState({
        email: '',
        password: '',
        isShop: false,
        firstName: '',
        lastName: '',
        shopName: '',
        phone: '',
        address: '',
        zip: '',
        city: '',
    });

    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setRegistrationData({
            ...registrationData,
            [name]: value
        });
    };

    const handleCheckboxChange = (event) => {
        const { name, checked } = event.target;
        setRegistrationData({
            ...registrationData,
            [name]: checked
        });
    };

    const handleLoginClick = () => {
        navigateTo('/login');
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const response = await fetch('http://localhost:9192/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(registrationData)
            });
            if (response.ok) {
                console.log('Registration successful');
                navigateTo('/login');
            } else {
                console.error('Registration failed:', response.statusText);
            }
        } catch (error) {
            console.error('Error registering:', error);
        }
    };

    return (
        <Grid container justifyContent="center" alignItems="center" spacing={2}>
            <Grid item xs={12} sm={8} md={6} lg={4}>
                <Typography variant="h2">Inscription d'un utilisateur</Typography>
                <form onSubmit={handleSubmit}>
                    <TextField
                        id="email"
                        name="email"
                        label="Email"
                        type="email"
                        value={registrationData.email}
                        onChange={handleInputChange}
                        fullWidth
                        required
                        margin="normal"
                    />
                    <TextField
                        id="password"
                        name="password"
                        label="Mot de passe"
                        type="password"
                        value={registrationData.password}
                        onChange={handleInputChange}
                        fullWidth
                        required
                        margin="normal"
                    />
                    {!registrationData.isShop && (
                        <>
                            <TextField
                                id="firstName"
                                name="firstName"
                                label="Prénom"
                                value={registrationData.firstName}
                                onChange={handleInputChange}
                                fullWidth
                                required
                                margin="normal"
                            />
                            <TextField
                                id="lastName"
                                name="lastName"
                                label="Nom"
                                value={registrationData.lastName}
                                onChange={handleInputChange}
                                fullWidth
                                required
                                margin="normal"
                            />
                        </>
                    )}
                    {registrationData.isShop && (
                        <>
                            <TextField
                                id="shopName"
                                name="shopName"
                                label="Nom du magasin"
                                value={registrationData.shopName}
                                onChange={handleInputChange}
                                fullWidth
                                required
                                margin="normal"
                            />
                            <TextField
                                id="phone"
                                name="phone"
                                label="Téléphone"
                                value={registrationData.phone}
                                onChange={handleInputChange}
                                fullWidth
                                required
                                margin="normal"
                            />
                            <TextField
                                id="address"
                                name="address"
                                label="Adresse"
                                value={registrationData.address}
                                onChange={handleInputChange}
                                fullWidth
                                required
                                margin="normal"
                            />
                            <TextField
                                id="zip"
                                name="zip"
                                label="Code Postal"
                                value={registrationData.zip}
                                onChange={handleInputChange}
                                fullWidth
                                required
                                margin="normal"
                            />
                            <TextField
                                id="city"
                                name="city"
                                label="Ville"
                                value={registrationData.city}
                                onChange={handleInputChange}
                                fullWidth
                                required
                                margin="normal"
                            />
                        </>
                    )}
                    <FormControlLabel
                        control={<Checkbox name="isShop" checked={registrationData.isShop} onChange={handleCheckboxChange} />}
                        label="S'inscrire en tant que magasin"
                    />
                    <div style={{ marginTop: '10px', display: 'flex', justifyContent: 'space-between' }}>
                        <Button type="button" variant="contained" style={{ backgroundColor: 'white', border: '2px solid #1976d2', color: '#1976d2' }} onClick={handleLoginClick}>
                            Se connecter
                        </Button>
                        <Button type="submit" variant="contained" color="primary">
                            S'inscrire
                        </Button>
                    </div>
                </form>
            </Grid>
        </Grid>
    );
};

export default Register;
