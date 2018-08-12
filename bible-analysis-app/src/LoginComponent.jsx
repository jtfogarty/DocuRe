import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Input from '@material-ui/core/Input';
import Button from '@material-ui/core/Button';

const styles = theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing.unit * 2,
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  input: {
    margin: theme.spacing.unit,
  },
});

//function CenteredGrid(props) {
class LoginComponent extends React.Component {
    //const { classes } = props;

  constructor(props) {
      super(props);
  }

  usernameChanged = e => {
      console.log('e', e);
      console.log('e.target.value: ', e.target.value);
      this.props.usernameChanged(e.target.value);
  }

  passwordChanged = e => {
    console.log('e', e);
    console.log('e.target.value: ', e.target.value);
    this.props.passwordChanged(e.target.value);
  }

  render() {
    const { classes } = this.props;    
    return (
        <div className={classes.root}>
        <Grid container spacing={24}>
            <Grid item xs={12}>
            <Paper className={classes.paper}>
                <h2>Welcome to the Document Research Application!</h2><br/>
            </Paper>
            <Paper className={classes.paper}>
                <h3>Username: </h3>
                <Input
                    defaultValue="Your username here"
                    className={classes.input}
                    inputProps={{
                    'aria-label': 'Description',
                    }}
                    onChange={this.usernameChanged}
                />
                <br/>
                <h3>Password: </h3>
                <Input
                    defaultValue="Your password here"
                    className={classes.input}
                    inputProps={{
                    'aria-label': 'Description',
                    }}
                    onChange={this.passwordChanged}
                />
                <br />
                <Button variant="outlined" className={classes.button} onClick={this.props.login}>
                    Login
                </Button>
            </Paper>
            </Grid>
        </Grid>
        </div>
    );
    }
}

LoginComponent.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(LoginComponent);
