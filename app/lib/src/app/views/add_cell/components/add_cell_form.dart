import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/components/default_button_component.dart';
import 'package:projeto_igreja/src/app/components/form_error_component.dart';
import 'package:projeto_igreja/src/app/components/theme_data.dart';
import 'package:projeto_igreja/src/app/constants.dart';
import 'package:projeto_igreja/src/app/models/cell.dart';
import 'package:projeto_igreja/src/app/provider/cells_provider.dart';
import 'package:projeto_igreja/src/app/size_config.dart';
import 'package:projeto_igreja/src/app/views/cellular_view/cellular_view.dart';
import 'package:projeto_igreja/src/app/views/sign_In/components/custom_svg_icon.dart';
import 'package:provider/provider.dart';

class AddCellForm extends StatefulWidget {
  @override
  _AddCellFormState createState() => _AddCellFormState();
}

class _AddCellFormState extends State<AddCellForm> {
  final _formKey = GlobalKey<FormState>();
  final List<String> errors = [];
  final Map<String, Object> _formData = {};

  String name;
  String address;
  String number;
  String date;
  TimeOfDay time;

  List listDate = [
    'Segunda-feira',
    'Terça-feira',
    'Quarta-feira',
    'Quinta-feira',
    'Sexta-feira',
    'Sábado',
    'Domingo',
  ];

  void addError({String error}) {
    if (!errors.contains(error)) {
      setState(() {
        errors.add(error);
      });
    }
  }

  void removeError({String error}) {
    if (errors.contains(error)) {
      setState(() {
        errors.remove(error);
      });
    }
  }

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    time = TimeOfDay.now();
    _formData['time'] = time.toString().substring(10, 15);
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: SingleChildScrollView(
        child: Column(
          children: [
            buildNameFormField(),
            SizedBox(height: getProportionateScreenHeight(30)),
            buildAdressNumber(),
            SizedBox(height: getProportionateScreenHeight(30)),
            buildDateTime(),
            SizedBox(height: getProportionateScreenHeight(10)),
            FormError(errors: errors),
            SizedBox(height: getProportionateScreenHeight(40)),
            DefaultButton(
              text: 'Salvar',
              press: () {
                if (_formKey.currentState.validate()) {
                  _formKey.currentState.save();

                  Provider.of<CellsProvider>(context, listen: false).put(
                    Cell(
                      id: _formData['id'],
                      name: _formData['name'],
                      address: _formData['address'],
                      number: _formData['number'],
                      date: _formData['date'],
                      time: _formData['time'],
                    ),
                  );
                  //Vai para Home
                  Navigator.pushReplacementNamed(
                      context, CellularView.routeName);
                }
              },
            ),
          ],
        ),
      ),
    );
  }

  TextFormField buildNameFormField() {
    return TextFormField(
      onSaved: (newValue) => _formData['name'] = newValue,
      onChanged: (value) {
        if (value.isNotEmpty) {
          removeError(error: kNamelNullError);
        }
      },
      validator: (value) {
        if (value.isEmpty) {
          addError(error: kNamelNullError);
          return "";
        }
        return null;
      },
      decoration: InputDecoration(
        labelText: 'Nome',
        hintText: 'Digite o Nome da  Célula',
        floatingLabelBehavior: FloatingLabelBehavior.always,
        suffixIcon: CustomSuffixIcon(
          svgIcon: 'assets/icons/hexagon.svg',
        ),
      ),
    );
  }

  Row buildAdressNumber() {
    return Row(
      mainAxisSize: MainAxisSize.min,
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      crossAxisAlignment: CrossAxisAlignment.baseline,
      children: [
        Expanded(
          flex: 2,
          child: TextFormField(
            onSaved: (newValue) => _formData['address'] = newValue,
            onChanged: (value) {
              if (value.isNotEmpty) {
                removeError(error: kNamelNullError);
              }
            },
            validator: (value) {
              if (value.isEmpty) {
                addError(error: kNamelNullError);
                return "";
              }
              return null;
            },
            decoration: InputDecoration(
              labelText: 'Endereço',
              hintText: 'Endereço',
              contentPadding:
                  new EdgeInsets.symmetric(vertical: 10, horizontal: 20),
              floatingLabelBehavior: FloatingLabelBehavior.always,
              suffixIcon: CustomSuffixIcon(
                svgIcon: 'assets/icons/hexagon.svg',
              ),
            ),
          ),
        ),
        SizedBox(width: 10),
        Expanded(
          flex: 1,
          child: TextFormField(
            keyboardType: TextInputType.number,
            onSaved: (newValue) => _formData['number'] = newValue,
            onChanged: (value) {
              if (value.isNotEmpty) {
                removeError(error: kPhoneNumberNullError);
              }
            },
            validator: (value) {
              if (value.isEmpty) {
                addError(error: kPhoneNumberNullError);
                return "";
              }
              return null;
            },
            decoration: InputDecoration(
              labelText: 'Número',
              alignLabelWithHint: true,
              contentPadding:
                  new EdgeInsets.symmetric(vertical: 20, horizontal: 20),
              floatingLabelBehavior: FloatingLabelBehavior.always,
            ),
          ),
        ),
      ],
    );
  }

  Row buildDateTime() {
    return Row(
      mainAxisSize: MainAxisSize.min,
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        //Dia
        Expanded(
          child: DropdownButtonFormField(
            onSaved: (newValue) => _formData['date'] = newValue,
            icon: Icon(Icons.arrow_drop_down_rounded),
            iconSize: getProportionateScreenHeight(30),
            iconDisabledColor: kPrimaryColor,
            iconEnabledColor: kPrimaryColor,

            style: TextStyle(
              color: kSecondaryColor,
              fontSize: getProportionateScreenHeight(18),
            ),

            value: _formData['date'],

            items: listDate.map((valueItem) {
              return DropdownMenuItem(
                value: valueItem,
                child: Text(valueItem, overflow: TextOverflow.visible),
              );
            }).toList(), //items

            hint: Text('Dia'),

            onChanged: (value) {
              if (value != 'Dia da Semana') {
                removeError(error: kCivilNullError);
              }
            },
            validator: (value) {
              if (value == 'Dia da Semana') {
                addError(error: kCivilNullError);
                return "";
              }
              return null;
            },

            decoration: inputDecorationRow(labelText: 'Dia da Semana'),
          ),
        ),
        SizedBox(width: getProportionateScreenWidth(20)),
        //Hora
        ElevatedButton(
          style: ElevatedButton.styleFrom(
            primary: Colors.white,
            side: BorderSide(color: kSecondaryColor),
            elevation: 0.0,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(28),
            ),
          ),
          onPressed: () {
            _pickTime();
          },
          child: Container(
            alignment: Alignment.center,
            height: 50.0,
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: <Widget>[
                Row(
                  children: <Widget>[
                    Container(
                      child: Row(
                        children: <Widget>[
                          Icon(
                            Icons.date_range,
                            size: 18.0,
                            color: kPrimaryColor,
                          ),
                          Text(
                            "${_formData['time']}",
                            style: TextStyle(
                                color: kSecondaryColor, fontSize: 18.0),
                          ),
                        ],
                      ),
                    )
                  ],
                ),
                Text(
                  " Hora",
                  style: TextStyle(color: kSecondaryColor, fontSize: 18.0),
                ),
              ],
            ),
          ),
        )
      ],
    );
  }

  _pickTime() async {
    TimeOfDay t = await showTimePicker(
        context: context,
        initialTime: time,
        builder: (BuildContext context, Widget child) {
          return MediaQuery(
            data: MediaQuery.of(context).copyWith(alwaysUse24HourFormat: true),
            child: child,
          );
        });
    if (t != null)
      setState(() {
        _formData['time'] = t.toString().substring(10, 15);
        print(t.minute.floor());
      });
  }
}
