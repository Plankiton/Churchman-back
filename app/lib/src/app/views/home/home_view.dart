import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:projeto_igreja/src/app/constants.dart';
import 'package:projeto_igreja/src/app/views/cellular_view/cellular_view.dart';
import 'package:projeto_igreja/src/app/views/church/church_view.dart';
import 'package:projeto_igreja/src/app/views/profile/profile_view.dart';
import '../../size_config.dart';
import 'components/home_body.dart';

class HomeView extends StatelessWidget {
  static String routeName = '/home';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBody: true,
      extendBodyBehindAppBar: true,
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(180),
        child: AppBar(
          elevation: 0,
          backgroundColor: Colors.transparent,
          shape: RoundedRectangleBorder(),
          brightness: Brightness.light,
          centerTitle: true,
          title: Text('FILADÉLFIA FORTALEZA',
              style: TextStyle(
                  color: Colors.white,
                  fontWeight: FontWeight.bold,
                  fontStyle: FontStyle.italic)),
          flexibleSpace: Container(
            padding: EdgeInsets.only(top: 55, left: 30, right: 30),
            height: MediaQuery.of(context).size.height,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.only(
                bottomLeft: Radius.elliptical(190, 80),
                bottomRight: Radius.elliptical(190, 80),
              ),
              gradient: kPrimaryGradientColor,
            ),
            child: Padding(
              padding: const EdgeInsets.only(top: 50),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      IconButton(
                        icon: SvgPicture.asset(
                          'assets/icons/church.svg',
                          height: getProportionateScreenWidth(100),
                          width: getProportionateScreenWidth(100),
                          color: Colors.white,
                        ),
                        iconSize: 50,
                        onPressed: () {
                          Navigator.pushNamed(context, ChurchView.routeName);
                        },
                      ),
                      IconButton(
                        icon: SvgPicture.asset(
                          'assets/icons/hexagon.svg',
                          height: getProportionateScreenWidth(100),
                          width: getProportionateScreenWidth(100),
                          color: Colors.white,
                        ),
                        iconSize: 50,
                        onPressed: () {
                          Navigator.pushNamed(context, CellularView.routeName);
                        },
                      ),
                      IconButton(
                        icon: SvgPicture.asset(
                          'assets/icons/User Icon.svg',
                          height: getProportionateScreenWidth(100),
                          width: getProportionateScreenWidth(100),
                          color: Colors.white,
                        ),
                        iconSize: 50,
                        onPressed: () {
                          Navigator.pushNamed(context, ProfileView.routeName);
                        },
                      ),
                    ],
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Padding(
                        padding: const EdgeInsets.only(left: 8),
                        child: Text('IGREJA',
                            style: TextStyle(
                              color: Colors.white,
                              fontWeight: FontWeight.bold,
                              fontStyle: FontStyle.italic,
                              height: 0.9,
                            )),
                      ),
                      Text(
                        'VISÃO\nCELULAR',
                        style: TextStyle(
                          color: Colors.white,
                          fontWeight: FontWeight.bold,
                          fontStyle: FontStyle.italic,
                          height: 0.9,
                        ),
                        textAlign: TextAlign.center,
                      ),
                      Padding(
                        padding: const EdgeInsets.only(right: 8),
                        child: Text(
                          'MEUS\nDADOS',
                          style: TextStyle(
                            color: Colors.white,
                            fontWeight: FontWeight.bold,
                            fontStyle: FontStyle.italic,
                            height: 0.9,
                          ),
                          textAlign: TextAlign.center,
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
      body: HomeBody(),
    );
  }
}
