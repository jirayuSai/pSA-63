import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import moment from 'moment';
import { EntPrescription } from '../../api/models/EntPrescription';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [Prescriptions, setPrescriptions] = useState<EntPrescription[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getPrescriptions = async () => {
     const res = await api.listPrescription({ limit: 10, offset: 0 });
     setLoading(false);
     setPrescriptions(res);
     console.log(res);
   };
   getPrescriptions();
 }, [loading]);

const deletePrescription = async (id: number) => {
 const res = await api.deletePrescription({ id: id });
 setLoading(true);
};
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">รหัสใบบันทึกการจ่ายยา</TableCell>
           <TableCell align="center">ชื่อแพทย์</TableCell>
           <TableCell align="center">ชื่อผู้ป่วย</TableCell>
           <TableCell align="center">รหัสยา</TableCell>
           <TableCell align="center">ชื่อเภสัชกร</TableCell>
           <TableCell align="center">date</TableCell> 
         </TableRow>
       </TableHead>
       <TableBody>
         {Prescriptions.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
             <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
             <TableCell align="center">{item.edges?.mmedicine?.id}</TableCell>
             <TableCell align="center">{item.edges?.systemmember?.systemmemberName}</TableCell>
             <TableCell align="center">{moment(item.datetime).format('DD/MM/YYYY  HH:mm')}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                  deletePrescription(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 DELETE
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
 
