import { createTheme, ThemeOptions } from '@material-ui/core';
import { PaletteOptions } from '@material-ui/core/styles/createPalette';

export const PALETTE = {
  type: 'dark',
  background: {
    default: 'rgb(17, 18, 23)',
    paper: '#191c23',
  },
  primary: {
    light: '',
    main: '#E31D38',
    dark: '#A20F23',
    contrastText: '',
  },
  text: {
    primary: '#FFFFFF',
    secondary: '#FEFEFE ',
  },
  grey: {
    
  },
  success: {
    main: '#38FF70',
    dark: '#155724',
    light: '#d4edda'
  }
}

const BREAKPOINTS = {
  values: {
    xs: 0,
    sm: 480,
    md: 780,
    lg: 1096,
    xl: 1200,
  }
}

const mainTheme = createTheme({
  spacing: 8,
  palette: PALETTE as PaletteOptions,
  breakpoints: BREAKPOINTS,
  overrides: {
    MuiTypography: {
      h1: {
        fontSize: 40,
        lineHeight: '48px',
        fontWeight: 900,
        marginBottom: 32,
      },
      h2: {
        textTransform: 'uppercase',
        fontSize: 28,
        fontWeight: 400,
        lineHeight: '33.6px',
        marginBottom: '14px',
        marginTop: 48,
      },
      body1: {
        fontSize: 16,
        fontWeight: 300,
        lineHeight: '24px',
      },
      body2: {
        fontSize: 14,
        fontWeight: 700,
        lineHeight: '17px',
      }
    },
    MuiInputBase: {
      root: {
        fontSize: 16,
        lineHeight: '24px',
        color: '#000',
        background: PALETTE.background.paper,
        border: '1px solid #919191',
        height: 40,
        borderRadius: 5,
        margin: ' 0',
      },
      input: {
        padding: 0,
        height: 40,
        backgroundColor: '#fff',
        paddingLeft: 16,
        borderRadius: '0 5px 5px 0',
      },
    },
    MuiInputAdornment: {
      root: {
        width: 48,
      },
      positionStart: {
        marginRight: 0,
        '& svg': {
          margin: '0 auto',
        },
      },
    },
    MuiButton: {
      root: {
        textTransform: 'none',
        borderRadius: 5,
        fontSize: 16,
        lineHeight: 1.5,
        whiteSpace: 'nowrap',
      },
      contained: {
        backgroundColor: PALETTE.primary.main,
        color: PALETTE.text.primary,
        fontWeight: 400,
        '&&:hover': {
          backgroundColor: PALETTE.primary.main,
        }
      },
      outlined: {
        borderColor: PALETTE.primary.main,
        backgroundColor: PALETTE.background.default,
        '&&:hover': {
          backgroundColor: PALETTE.background.paper,
        }
      },
      text: {
        backgroundColor: PALETTE.background.default,
        color: PALETTE.text.primary,
        fontWeight: 300,
        '&&:hover': {
          backgroundColor: PALETTE.background.default,
          color: PALETTE.primary.main,
        }
      }
    },
  }
} as ThemeOptions )

export { mainTheme };