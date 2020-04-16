import React, {Component} from 'react';
import {View, Text, StatusBar, StyleSheet} from 'react-native';

class HomeScreeen_Parent extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <>
        <StatusBar hidden={true} />
        <View style={styles.screen}>
          <Text style={styles.title}>Home Screen Parent</Text>
        </View>
      </>
    );
  }
}
const styles = StyleSheet.create({
  screen: {
    flex: 1,
    backgroundColor: 'white',
    color: 'black',
    justifyContent: 'center',
    alignItems: 'center',
  },
  title: {
    fontSize: 40,
    color: 'skyblue',
  },
});
export default HomeScreeen_Parent;
