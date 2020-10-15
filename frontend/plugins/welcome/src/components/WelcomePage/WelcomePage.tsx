import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'แผนกจ่ายยา' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ยินดีต้อนรับ ${profile.givenName || 'to Backstage'}`}
       subtitle="Some quick intro and links."
     ></Header>
     <Content>
       <ContentHeader title="ประวัติการจ่ายยา">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             Add
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
 
