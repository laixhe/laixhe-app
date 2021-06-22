import 'package:flutter/material.dart';

import 'chat_tag/chat_tag.dart';
import 'friend_tag/friend_tag.dart';
import 'my_tag/my_tag.dart';

class MainNavi extends StatefulWidget {
  const MainNavi({Key? key}) : super(key: key);

  @override
  _MainNaviState createState() => _MainNaviState();
}

class _MainNaviState extends State<MainNavi> {

  final List<Widget> _children = [
    ChatTag(),
    FriendTag(),
    MyTag()
  ];

  int _currentIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _children[_currentIndex],
      bottomNavigationBar: BottomNavigationBar(
        selectedItemColor: Colors.black,  // 选中颜色
        unselectedItemColor: Colors.grey, // 未选中颜色
        type: BottomNavigationBarType.fixed, // 未选中时的导航-显示文本和图标
        //type: BottomNavigationBarType.shifting, // 未选中时的导航-只显示图标
        selectedFontSize: 12.0,   // 选中字体大小 (默认选中的比不选中大点)
        unselectedFontSize: 12.0, // 未选中字体大小

        currentIndex: _currentIndex, // 绑定当前选中的下标
        items: <BottomNavigationBarItem>[
          BottomNavigationBarItem(
              icon: Icon(Icons.contact_mail),
              label: '聊天'
          ),
          BottomNavigationBarItem(
              icon: Icon(Icons.contacts),
              label: '好友'
          ),
          BottomNavigationBarItem(
              icon: Icon(Icons.person),
              label: '我的'
          ),
        ],
        onTap: (int selectIndex){
          // 点击事件
          setState(() {
            _currentIndex = selectIndex;
          });
        },
      ),
    );
  }

}
