import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  Link,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
//import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import MenuItem from '@material-ui/core/MenuItem';
//import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';


import { EntDoctor, EntMedicine, EntPatient, EntPrescription, EntSystemmember } from '../../api';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);
export default function CreatePrescription() {
  const classes = useStyles();
  const profile = { givenName: ''};
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [doctors, setDoctors] = useState<EntDoctor[]>([]);
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [medicines, setMedicines] = useState<EntMedicine[]>([]);
  const [systemmembers, setSysemmembers] = useState<EntSystemmember[]>([]);
  const [prescription, setPrescriptions] = useState<EntPrescription[]>([]);

  const [loading, setLoading] = useState(true);

  const [doctorID, setDoctor] = useState(Number);
  const [patientID, setPatient] = useState(Number);
  const [medicineID, setMedicine] = useState(Number);
  const [systemmemberID, setSysemmember] = useState(Number);
  const [prescriptionID, setPrescription] = useState(Number);
  const [datetime, setDateTime] = useState(String);

  useEffect(() => {
    const getDoctors = async () => {
      const res = await api.listDoctor({ limit: 10, offset: 0 });
      setLoading(false);
      setDoctors(res);
    };
    getDoctors();

    const getPatients = async () => {
      const res = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatients(res);
      console.log(res);
    };
    getPatients();

    const getMedicines = async () => {
      const res = await api.listMedicine({ limit: 10, offset: 0 });
      setLoading(false);
      setMedicines(res);
    };
    getMedicines();

    const getSystemmembers = async () => {
      const res = await api.listSystemmember({ limit: 10, offset: 0 });
      setLoading(false);
      setSysemmembers(res);
    };
    getSystemmembers();

  }, [loading]);

  const DoctorhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDoctor(event.target.value as number);
  };

  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatient(event.target.value as number);
  };

  const MedicinehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setMedicine(event.target.value as number);
  };

  const SystemmemberhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSysemmember(event.target.value as number);
  };
  const DateTimehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDateTime(event.target.value as string);
  };


  const CreatePrescription = async () => {
    const prescription = {
      doctor: doctorID,
      patient: patientID,
      systemmember: systemmemberID,
      datetime: datetime + ":00+07:00",

    };
    console.log(prescription);
    const res: any = await api.createPrescription({ prescription: prescription });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
      title={`${profile.givenName || 'Prescription'}`}
      subtitle="BLUE MOON DORMITORY"
     ></Header>
      <Content>
        <ContentHeader title="ADD DATA">
          <div>
            <Link component={RouterLink} to="/">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              Exit
            </Button>
          </Link>
          </div>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  success!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          <FormControl
                className={classes.margin}
                variant="outlined"
              >
              <Typography variant="h6" gutterBottom  align="center">
                Doctor Name : 
                <Typography variant="body1" gutterBottom> 
                <Select
                  labelId="doctor"
                  id="doctor"
                  value={doctorID}
                  onChange={DoctorhandleChange}
                  style={{ width: 400 }}
                >
                {doctors.map((item: EntDoctor) => (
                  <MenuItem value={item.id}>{item.doctorName}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography>
              </FormControl>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <Typography variant="h6" gutterBottom  align="center">
                Patient Name : 
                <Typography variant="body1" gutterBottom> 
                <Select
                  labelId="patient"
                  id="patient"
                  value={patientID}
                  onChange={PatienthandleChange}
                  style={{ width: 400 }}
                >
                {patients.map((item: any) => (
                  <MenuItem value={item.id}>{item.patientName}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography>
              </FormControl>
            </div>

            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography variant="h6" gutterBottom  align="center">
                Medicine ID : 
                <Typography variant="body1" gutterBottom> 
              <Select
                labelId="medicine"
                id="medicine"
                value={medicineID}
                onChange={MedicinehandleChange}
                style={{ width: 400 }}
              >
                {medicines.map((item: EntMedicine) => (
                  <MenuItem value={item.id}>{item.id}</MenuItem>
                ))}
              </Select>
                </Typography>
                </Typography>
            </FormControl>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <Typography variant="h6" gutterBottom  align="center">
                  Systemmember Name : 
                <Typography variant="body1" gutterBottom> 
                <Select
                   labelId="systemmember"
                   id="systemmember"
                   value={systemmemberID}
                   onChange={SystemmemberhandleChange}
                   style={{ width: 400 }}
                >
                  {systemmembers.map((item: EntSystemmember) => (
                  <MenuItem value={item.id}>{item.systemmemberName}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography>
              </FormControl>
            </div>
            <tr>
            <td>
            <FormControl
                  fullWidth
                  className={classes.margin}
                  variant="outlined"
                >
                  <form className={classes.root} noValidate>
                      <TextField
                        id="datetime-local"
                        label="Date"
                        type="datetime-local"
                        //defaultValue="2017-05-24T10:30"
                        className={classes.textField}
                        onChange={DateTimehandleChange}
                        InputLabelProps={{
                          shrink: true,
                        }}
                      />
                </form>
              </FormControl>
            </td>
          </tr>

            <div className={classes.margin}>
            <Typography variant="h6" gutterBottom  align="center">
              <Button
                onClick={() => {
                  CreatePrescription();
                }}
                variant="contained"
                color="primary"
              >
                Submit
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                Back
             </Button>
              </Typography>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}