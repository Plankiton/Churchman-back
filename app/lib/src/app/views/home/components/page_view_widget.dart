import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/provider/events_provider.dart';
import 'package:projeto_igreja/src/app/views/home/components/slide_tile.dart';
import 'package:provider/provider.dart';

class PageViewWidget extends StatefulWidget {
  @override
  _PageViewWidgetState createState() => _PageViewWidgetState();
}

class _PageViewWidgetState extends State<PageViewWidget> {
  int _currentPage = 0;

  final PageController _pageController = PageController(viewportFraction: 0.9);

  @override
  void initState() {
    _pageController.addListener(() {
      int next = _pageController.page.round();
      if (_currentPage != next) {
        setState(() {
          _currentPage = next;
        });
      }
    });
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final EventsProvider events = Provider.of(context);

    return Stack(children: [
      PageView.builder(
        itemCount: events.count,
        controller: _pageController,
        itemBuilder: (context, index) {
          bool activePage = index == _currentPage;
          return SlideTile(
            events: events,
            index: index,
            activePage: activePage,
          );
        },
      ),
    ]);
  }
}
