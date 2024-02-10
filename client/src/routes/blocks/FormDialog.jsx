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
import { TimePicker } from '@mui/x-date-pickers/TimePicker';
import dayjs from 'dayjs';

export default function FormDialog({ shopDetails }) {
    const [open, setOpen] = React.useState(false);
    const [resa, setResa] = React.useState(
        {
            date: "",
            hours: "",
            employee_id: null,
            user_id: "3",
            shop_id: shopDetails?.user._id
        }
    )

    console.log(resa)
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

    const handleTimeChange = (time) => {
        const formattedTime = dayjs(time).format('HH:mm:ss');
        setResa({ ...resa, hours: formattedTime });
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
                        try {
                            const requestOptions = {
                                method: 'POST',
                                headers: { 'Content-Type': 'application/json' },
                                body: JSON.stringify(resa)
                            };

                            const response = await fetch('http://localhost:9192/createReservation', requestOptions);

                            if (!response.ok) {
                                throw new Error('Failed to create reservation');
                            }

                        } catch (error) {
                            console.error('Error creating reservation:', error);
                        }
                        handleClose();
                    },
                }}
            >
                <DialogTitle>Prendre un rendez-vous</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        To subscribe to this website, please enter your email address here. We
                        will send updates occasionally.
                    </DialogContentText>
                    <DatePicker onChange={handleDateChange} />
                    <TimePicker onChange={handleTimeChange} />
                    <FormControl fullWidth>
                        <InputLabel id="demo-simple-select-label">Employee</InputLabel>
                        <Select
                            value={resa.employee}
                            label="Employee"
                            onChange={handleSelectChange}
                        >
                            {shopDetails && shopDetails.hairdressers.map((e) => {
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