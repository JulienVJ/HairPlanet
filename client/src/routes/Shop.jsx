import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import Container from '@mui/material/Container';
import { Avatar, Box, ImageList, ImageListItem, List, ListItem, ListItemAvatar, ListItemText, Typography } from '@mui/material';
import HomeIcon from '@mui/icons-material/Home';
import PermContactCalendarIcon from '@mui/icons-material/PermContactCalendar';
import PeopleIcon from '@mui/icons-material/People';
import FormDialog from './blocks/FormDialog'
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'

const itemData = [
    {
        img: "https://plus.unsplash.com/premium_photo-1677616799911-786522e9a1d1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OXx8SGFpciUyMGRyZXNzZXJ8ZW58MHx8MHx8fDA%3D",
    },
    {
        img: "https://plus.unsplash.com/premium_photo-1677616799911-786522e9a1d1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OXx8SGFpciUyMGRyZXNzZXJ8ZW58MHx8MHx8fDA%3D",
    },
    {
        img: "https://plus.unsplash.com/premium_photo-1677616799911-786522e9a1d1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OXx8SGFpciUyMGRyZXNzZXJ8ZW58MHx8MHx8fDA%3D",
    },
    {
        img: "https://plus.unsplash.com/premium_photo-1677616799911-786522e9a1d1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OXx8SGFpciUyMGRyZXNzZXJ8ZW58MHx8MHx8fDA%3D",
    },
    {
        img: "https://plus.unsplash.com/premium_photo-1677616799911-786522e9a1d1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OXx8SGFpciUyMGRyZXNzZXJ8ZW58MHx8MHx8fDA%3D",
    },
    {
        img: "https://plus.unsplash.com/premium_photo-1677616799911-786522e9a1d1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OXx8SGFpciUyMGRyZXNzZXJ8ZW58MHx8MHx8fDA%3D",
    },
]

function Shop() {
    const { shopName } = useParams();
    const [shopDetails, setShopDetails] = useState()

    useEffect(() => {
        const fetchShopDetails = async (shopName) => {
            try {
                const response = await fetch(`http://localhost:9192/shopDetails?shopName=${shopName}`);
                const data = await response.json();
                setShopDetails(data)
            } catch (error) {
                console.error('Error fetching shop details:', error);
            }
        }

        fetchShopDetails(shopName);
    }, []);

    return (
        <Container>
            <Box>
                <h1>{shopDetails?.user.shopName}</h1>
                <Box sx={{ display: "flex", padding: 3 }}>
                    <ImageList
                        sx={{ width: 500 }}
                        variant="quilted"
                        cols={3}
                        rowHeight={150}
                    >
                        {itemData.map((item) => (
                            <ImageListItem key={item.img} cols={item.cols || 1} rows={item.rows || 1}>
                                <img
                                    srcSet={`${item.img}?w=164&h=164&fit=crop&auto=format&dpr=2 2x`}
                                    src={`${item.img}?w=164&h=164&fit=crop&auto=format`}
                                    alt={item.title}
                                    loading="lazy"
                                />
                            </ImageListItem>
                        ))}
                    </ImageList>
                    <List sx={{ width: '100%', maxWidth: 360 }}>
                        <ListItem>
                            <ListItemAvatar>
                                <Avatar>
                                    <PermContactCalendarIcon />
                                </Avatar>
                            </ListItemAvatar>
                            <ListItemText primary="Contact" secondary={shopDetails?.user.phone} />
                        </ListItem>
                        <ListItem>
                            <ListItemAvatar>
                                <Avatar>
                                    <HomeIcon />
                                </Avatar>
                            </ListItemAvatar>
                            <ListItemText primary="Adresse" secondary={shopDetails?.user.address} />
                        </ListItem>
                        <ListItem>
                            <ListItemAvatar>
                                <Avatar>
                                    <PeopleIcon />
                                </Avatar>
                            </ListItemAvatar>
                            <ListItemText
                                primary="Coiffeurs"
                                secondary={
                                    (shopDetails?.hairdressers || []).map((h, index) => (
                                        <Typography key={index} component="span" variant="body2" color="textPrimary">
                                            {`${h?.FirstName} ${h?.LastName}`}
                                            <br />
                                        </Typography>
                                    ))
                                }
                            />
                        </ListItem>
                    </List>
                </Box>
            </Box>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
                <FormDialog shopDetails={shopDetails} />
            </LocalizationProvider>
        </Container>
    )
}

export default Shop;
