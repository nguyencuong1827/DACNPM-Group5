// In App.js in a new project

import * as React from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {createStackNavigator} from '@react-navigation/stack';
import LoginScreen from './LoginScreen';
import RegisterScreen from './RegisterScreen';
import HomeScreen_Parent from './HomeScreen_Parent';
import HomeScreen_Child from './HomeScreen_Child';

const Stack = createStackNavigator();

function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName="Login">
        <Stack.Screen
          name="Login"
          component={LoginScreen}
          options={{
            title: 'Đăng nhập',
            headerTitleAlign: 'center',
          }}
        />
        <Stack.Screen
          name="Register"
          component={RegisterScreen}
          options={{
            title: 'Đăng ký',
            headerTitleAlign: 'center',
          }}
        />
        <Stack.Screen
          name="HomeScreen_Parent"
          component={HomeScreen_Parent}
          options={{
            title: 'Home',
            headerTitleAlign: 'center',
          }}
        />
        <Stack.Screen
          name="HomeScreen_Child"
          component={HomeScreen_Child}
          options={{
            title: 'Home',
            headerTitleAlign: 'center',
          }}
        />
      </Stack.Navigator>
    </NavigationContainer>
  );
}

export default App;
