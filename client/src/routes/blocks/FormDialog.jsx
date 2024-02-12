import * as React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { FormControl, InputLabel, MenuItem, Select } from '@mui/material';
import PropTypes from 'prop-types'
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import dayjs from 'dayjs';
import emailjs from '@emailjs/browser';

const hours = [
    { label: '09:00', value: '0900' },
    { label: '09:30', value: '0930' },
    { label: '10:00', value: '1000' },
    { label: '10:30', value: '1030' },
    { label: '11:00', value: '1100' },
    { label: '11:30', value: '1130' },
    { label: '12:00', value: '1200' },
    { label: '12:30', value: '1230' },
    { label: '13:00', value: '1300' },
    { label: '13:30', value: '1330' },
    { label: '14:00', value: '1400' },
    { label: '14:30', value: '1430' },
    { label: '15:00', value: '1500' },
    { label: '15:30', value: '1530' },
    { label: '16:00', value: '1600' },
    { label: '16:30', value: '1630' },
    { label: '17:00', value: '1700' },
    { label: '17:30', value: '1730' },
    { label: '18:00', value: '1800' },
    { label: '18:30', value: '1830' },
    { label: '19:00', value: '1900' }
];

export default function FormDialog({ shopDetails }) {
    const [open, setOpen] = React.useState(false);
    const [resa, setResa] = React.useState(
        {
            date: "",
            hours: "",
            employee_id: null,
            user_id: null,
            shop_id: null
        }
    )
    React.useEffect(() => {
        if (shopDetails) {
            const storedUserId = localStorage.getItem('userId');
            setResa({
                ...resa,
                shop_id: shopDetails.user._id,
                user_id: storedUserId
            });
        }
    }, [shopDetails]);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleSelectChange = (event) => {
        setResa({ ...resa, employee_id: event.target.value });
    };
    const handleDateChange = (date) => {
        const formattedDate = dayjs(date).format('ddd MMM DD YYYY HH:mm:ss [GMT]Z');
        setResa({ ...resa, date: formattedDate });
    };

    return (
        <React.Fragment>
            <Button variant="contained" onClick={handleClickOpen}>
                Prendre un rendez-vous
            </Button>
            <Dialog
                open={open}
                onClose={handleClose}
                PaperProps={{
                    component: 'form',
                    onSubmit: async () => {
                        const requestOptions = {
                            method: 'POST',
                            headers: { 'Content-Type': 'application/json' },
                            body: JSON.stringify(resa)
                        };

                        const response = await fetch('http://localhost:9192/createReservation', requestOptions)
                            .catch(error => {
                                console.error('Network error:', error);
                            });

                        if (!response.ok) {
                            throw new Error('Failed to create reservation');
                        }
                        const content = {
                            shop_name: shopDetails.user.shopName,
                            user_name: "Mat",
                            user_mail: "mlj.bouillon@outlook.fr",
                            date: resa.date,
                            hours: resa.hours
                        }

                        emailjs
                            .sendForm('service_s18yxti', 'template_68vf4gp', content, {
                                publicKey: 'lf4vpbg6q-_JhSsIn',
                            })
                            .then(
                                () => {
                                    console.log('SUCCESS!');
                                },
                                (error) => {
                                    console.log('FAILED...', error.text);
                                },
                            );
                    },
                }}
            >
                <DialogTitle>Prendre un rendez-vous</DialogTitle>
                <DialogContent flexItem>
                    <DialogContentText>
                        To subscribe to this website, please enter your email address here. We
                        will send updates occasionally.
                    </DialogContentText>
                    <div>
                        <DatePicker onChange={handleDateChange} />
                    </div>
                    {resa.date && hours.map((h) => {
                        const isReserved = shopDetails.reservations.some((r) => r.date === resa.date && r.hours === h.value);
                        return (
                            <Button
                                key={h.value}
                                variant={resa.hours === h.value ? "contained" : "outlined"}
                                onClick={() => setResa({ ...resa, hours: h.value })}
                                disabled={isReserved}
                            >
                                {h.label}
                            </Button>
                        );
                    })}

                    <FormControl fullWidth>
                        <InputLabel id="demo-simple-select-label">Employee</InputLabel>
                        <Select
                            value={resa.employee}
                            label="Employee"
                            onChange={handleSelectChange}
                        >
                            {(shopDetails?.hairdressers || []).map((e) => {
                                const hairdressersName = `${e.FirstName} ${e.LastName}`;
                                return (
                                    <MenuItem key={e.id} value={e.ID}>{hairdressersName}</MenuItem>
                                );
                            })
                            }
                        </Select>
                    </FormControl>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>Annuler</Button>
                    <Button type="submit">Accepter</Button>
                </DialogActions>
            </Dialog>
        </React.Fragment>
    );
}

FormDialog.propTypes = {
    shopDetails: PropTypes.object.isRequired,
};