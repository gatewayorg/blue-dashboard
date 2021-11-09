import React from 'react';
import { 
  StyledComponentProps,
  StyleRules, 
  Theme, 
  withStyles,
  Box,
  Button
} from '@material-ui/core';
import { Field, Form, FormRenderProps } from "react-final-form";
import { client } from '../../api/service';
import { t } from '../../common/utils/intl';
import LOGINBG from '../../assets/logins.jpeg';
import { InputField } from "../../common/Input";
import { PasswordInputField } from "../../common/Input/PasswordInputField"
import { onFormSubmit } from "../../common/formTypes";
import { useHistory } from 'react-router';
import { PATH } from '../../nav/const';
import { useSnackbar } from 'notistack';

const styles = (theme: Theme): StyleRules => ({
  root:{
    minHeight:'500px',
    position:'relative',
  },
  bgImg:{
    position:'fixed',
    left:0,
    right:0,
    top:0,
    bottom:0,
    // overflow:'hidden',
    // minHeight:'500px'
  },
  imgs:{
    maxWidth:2000
  },
  content:{
    position:'absolute',
    width:'100%',
    height:'100%',
    top:0,
    left:0,
  },
  container:{
    position:'absolute',
    width:400,
    height:400,
    top:'50%',
    left:'50%',
    backgroundColor:'rgba(0,0,0,0.5)',
    borderRadius:'20px',
    marginTop:'-200px',
    marginLeft:'-200px',
    padding:'20px'
  },
  formRoot:{
    "& .MuiOutlinedInput-input":{
      padding:'0',
      backgroundColor:'rgba(0,0,0,0)',
      color:'#fff',
      textIndent:'16px',
    },
    "& .MuiInputBase-root":{
      margin:'0',
      border:'none',
    },
  },
  title:{
    textAlign:'center',
    fontSize:'32px',
    margin:'0 16px'
  },
  buttons:{
    width:"100%",
    backgroundColor:'#099639',
    color:'#fff',
    opacity:'0.8',
    height:'42px',
    fontWeight:'bold',
    marginTop:"16px",
    '&&:hover': {
      backgroundColor: '#099639',
      color:'#fff',
      opacity:'1'
    }
  }
});


interface ILoginProps
extends StyledComponentProps {

}

const LoginComponent = ({
  classes = {},
}: ILoginProps) => {
  const history = useHistory();
  const { enqueueSnackbar } = useSnackbar();
  const onSubmit = (value:any)=>{
    client.post(`/api/user/login`,{
      "username":value.name,
      "password":value.password
    }).then(res => {
      if (res.status === 200) {
        window.sessionStorage.setItem('token',res.data.access_token);
        window.sessionStorage.setItem('username',value.name);
        history.push(PATH.DASHBOARD)
      }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  };

  const userForm = ({ handleSubmit }: FormRenderProps) => {
    return (
      <form onSubmit={handleSubmit}>
        <Box m={-2} mt={2}  className={classes.formRoot}>
          <Box p={2}>
            <Box pb={1}>{t("login.usernames")}</Box>
            <Field name="name"  placeholder={t("login.username")} component={InputField} margin="none" />
          </Box>
          <Box p={2}>
            <Box pb={1}>{t("login.passwords")}</Box>
            <Field name="password" placeholder={t("login.password")} component={PasswordInputField} margin="none" />
          </Box>
          <Box p={2}>
              <Button
                type="submit"
                color="primary"
                className={classes.buttons}
              >
                 {t("login.btn")}
              </Button>
            </Box>
        </Box>
      </form>
    );
  };

  return (
    <div className={classes.root}>   
      <div className={classes.bgImg}>
        <img src={LOGINBG} alt="" className={classes.imgs}/>
      </div>
      <div className={classes.content}>
        <div className={classes.container}>
          <p className={classes.title}>{t("login.title")}</p>
          <Form
            render={userForm}
            onSubmit={onSubmit as onFormSubmit}
          />
        </div>
      </div>
    </div>
  )
}

const Login = withStyles(styles)(LoginComponent);
export { Login };