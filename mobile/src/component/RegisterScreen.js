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
import RadioForm, {
  RadioButton,
  RadioButtonInput,
  RadioButtonLabel,
} from 'react-native-simple-radio-button';

function RegisterScreen({navigation}) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [typeAccount, setTypeAccount] = useState('');
  const [valueIndex, setValueIndex] = useState(0);
  const listTypeAccount = [
    {label: 'Cha mẹ', value: 'Parent'},
    {label: 'Trẻ em', value: 'Child'},
  ];
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
            placeholder="Email ... "
            placeholderTextColor="white"
            onChangeText={(username) => setUsername(username)}
            defaultValue={username}
            keyboardType="email-address"
          />
          <TextInput
            style={styles.textInput}
            placeholder="Mật khẩu ... "
            placeholderTextColor="white"
            onChangeText={(password) => setPassword(password)}
            defaultValue={password}
            secureTextEntry={true}
            keyboardType="default"
          />
          <TextInput
            style={styles.textInput}
            placeholder="Xác nhận mật khẩu ... "
            placeholderTextColor="white"
            onChangeText={(confirmPassword) => setPassword(confirmPassword)}
            defaultValue={password}
            secureTextEntry={true}
            keyboardType="default"
          />
          <RadioForm
            style={styles.radioForm}
            formHorizontal={true}
            animation={true}>
            {/* To create radio buttons, loop through your array of options */}
            {listTypeAccount.map((obj, i) => (
              <RadioButton
                style={i === 1 ? styles.radioButton : null}
                labelHorizontal={false}
                key={i}>
                {/*  You can set RadioButtonLabel before RadioButtonInput */}
                <RadioButtonInput
                  obj={obj}
                  index={i}
                  isSelected={valueIndex === i}
                  onPress={(value) => {
                    setTypeAccount(value);
                    setValueIndex(i);
                  }}
                  borderWidth={1}
                  buttonInnerColor={'#e74c3c'}
                  buttonOuterColor={'black'}
                  buttonSize={25}
                  buttonOuterSize={35}
                  buttonStyle={{}}
                  buttonWrapStyle={{marginLeft: 10}}
                />
                <RadioButtonLabel
                  obj={obj}
                  index={i}
                  labelHorizontal={true}
                  onPress={(value) => {
                    setTypeAccount(value);
                    setValueIndex(i);
                  }}
                  labelStyle={{fontSize: 20, color: 'black'}}
                  labelWrapStyle={{}}
                />
              </RadioButton>
            ))}
          </RadioForm>
          <TouchableOpacity
            onPress={() => {
              if (
                username === '' ||
                password === '' ||
                confirmPassword === ''
              ) {
                ToastAndroid.show(
                  'Vui lòng điền đủ thông tin',
                  ToastAndroid.SHORT,
                );
                return;
              }
              if (password !== confirmPassword) {
                ToastAndroid.show(
                  'Xác nhận mật khẩu không đúng',
                  ToastAndroid.SHORT,
                );
              }
              if (username === 'parent@gmail.com') {
                ToastAndroid.show('Tài khoản đã tồn tại', ToastAndroid.SHORT);
              } else {
                ToastAndroid.show('Đăng ký thành công', ToastAndroid.SHORT);
                navigation.navigate('Login');
              }
            }}>
            <View style={styles.button}>
              <Text style={styles.textButton}>Đăng ký</Text>
            </View>
          </TouchableOpacity>
          <TouchableOpacity
            onPress={() => {
              navigation.navigate('Login');
            }}>
            <View style={styles.button}>
              <Text style={styles.textButton}>Quay lại</Text>
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
  radioForm: {
    marginTop: 10,
    marginBottom: 10,
    marginLeft: 80,
    marginRight: 80,
  },
  radioButton: {
    marginLeft: 100,
  },
});

export default RegisterScreen;
