import React, {Component, useState} from 'react';
import {
  View,
  TextInput,
  StatusBar,
  StyleSheet,
  Text,
  Image,
  ToastAndroid,
} from 'react-native';
import {TouchableOpacity} from 'react-native-gesture-handler';

function LoginScreen({navigation}) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  return (
    <>
      <StatusBar hidden={true} />
      <View style={styles.screen}>
        <View style={styles.title}>
          <Text style={styles.textTitle}>Child Safety</Text>
          <Image source={require('../asset/images/child.png')} />
        </View>
        <View style={styles.input}>
          <TextInput
            style={styles.textInput}
            placeholder="Email ..."
            placeholderTextColor="white"
            onChangeText={(username) => setUsername(username)}
            defaultValue={username}
            keyboardType="email-address"
          />
          <TextInput
            style={styles.textInput}
            placeholder="Mật khẩu ..."
            placeholderTextColor="white"
            onChangeText={(password) => setPassword(password)}
            defaultValue={password}
            secureTextEntry={true}
            keyboardType="default"
          />
          <TouchableOpacity
            onPress={() => {
              if (
                (username === 'parent@gmail.com' && password === 'parent') ||
                (username === 'parent2@gmail.com' && password === 'parent2')
              ) {
                navigation.navigate('HomeScreen_Parent');
              } else if (
                username === 'child@gmail.com' &&
                password === 'child'
              ) {
                navigation.navigate('HomeScreen_Child');
              } else {
                ToastAndroid.show(
                  'Tài khoản hoặc mật khẩu không đúng',
                  ToastAndroid.SHORT,
                );
              }
            }}>
            <View style={styles.button}>
              <Text style={styles.textButton}>Đăng nhập</Text>
            </View>
          </TouchableOpacity>
          <TouchableOpacity
            onPress={() => {
              navigation.navigate('Register');
            }}>
            <View style={styles.button}>
              <Text style={styles.textButton}>Tạo tài khoản</Text>
            </View>
          </TouchableOpacity>
        </View>
      </View>
    </>
  );
}

const styles = StyleSheet.create({
  screen: {
    flex: 1,
    backgroundColor: '#009688',
  },
  title: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  textTitle: {
    fontSize: 72,
    fontFamily: 'Linotte Bold',
    fontStyle: 'italic',
  },
  input: {
    flex: 2,
    justifyContent: 'flex-start',
    alignItems: 'stretch',
  },
  textInput: {
    height: 66,
    borderColor: 'white',
    borderBottomWidth: 2,
    fontSize: 20,
    paddingTop: 20,
    paddingBottom: 20,
    paddingLeft: 10,
    marginBottom: 10,
    marginLeft: 100,
    marginRight: 100,
  },
  button: {
    borderWidth: 1,
    borderRadius: 40,
    borderColor: '#000000',
    justifyContent: 'center',
    alignItems: 'center',
    paddingTop: 20,
    paddingBottom: 20,
    marginTop: 10,
    marginBottom: 10,
    marginLeft: 100,
    marginRight: 100,
  },
  textButton: {
    fontSize: 20,
  },
});

export default LoginScreen;
