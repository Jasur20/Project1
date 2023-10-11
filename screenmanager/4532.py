import os

from kivy.config import Config

from kivymd.uix.button import *
from kivymd.uix.dialog import MDDialog

Config.set("graphics", "resizable", 0)
Config.set("graphics", "width", 1520)
Config.set("graphics", "height", 850)
from kivymd.uix.screen import MDScreen
from kivymd.uix.screenmanager import MDScreenManager
from kivymd.uix.toolbar import MDTopAppBar
from kivymd.uix.navigationdrawer import MDNavigationLayout,\
    MDNavigationDrawer,MDNavigationDrawerItem\
    ,MDNavigationDrawerMenu,MDNavigationDrawerDivider,\
    MDNavigationDrawerHeader,MDNavigationDrawerLabel
from kivymd.uix.fitimage import FitImage
from os import *
from kivymd.uix.pickers import MDDatePicker
from kivymd.uix.behaviors import MagicBehavior
from kivy.lang import Builder

from kivymd.uix.navigationdrawer import *
from kivymd.uix.button import MDFlatButton,MDIconButton
from kivymd.uix.button import MDRaisedButton

from kivymd.app import MDApp
from kivymd.uix.textfield import MDTextField

from kivymd.uix.button import MDFloatingActionButtonSpeedDial
from kivy.uix.screenmanager import Screen,ScreenManager

from kivymd.uix.hero import MDHeroFrom,MDHeroTo
from kivymd.uix.label import MDLabel

from kivymd.uix.boxlayout import MDBoxLayout

from kivymd.uix.floatlayout import MDFloatLayout
from kivymd.uix.responsivelayout import MDResponsiveLayout
from kivymd.uix.navigationdrawer import MDNavigationLayout, MDNavigationDrawer



class MagicButton(MagicBehavior,MDRaisedButton):
    pass

class Main(Screen):

    def __init__(self,**kw):
        super(Main,self).__init__(**kw)

        self.ids.tx.password=True
        self.ids.tx.icon_left='key-variant'
        self.ids.icon1.theme_icon_color='Primary'

    def press_off(self):
        os.system('shutdown -p')
    def on_save(self, instance, value, date_range):


        print(instance, value, date_range)

    def on_cancel(self, instance, value):
        pass

    def show_data_picker(self):
        data_dialog=MDDatePicker(
            mode="range"
        )
        data_dialog.bind(on_save=self.on_save,on_cancel=self.on_cancel)
        data_dialog.open()

    def growing(self):
        self.ids.mbutton1.grow()


    def shaking(self):
        #a=FitlImage()
        self.ids.mbutton1.shake()


    def twisting(self):
        self.ids.mbutton1.twist()


    def password(self):


        if self.ids.tx.text=='1234':
            self.manager.current='up'
        else:
            self.ids.tx.helper_text="Error"
            self.shaking()

    def unknown(self):
        self.ids.icon1.icon="eye" if self.ids.icon1.icon=='eye-off' \
            else 'eye-off'
        self.ids.tx.password=False if self.ids.tx.password \
                                           is True else True

    def change_password(self,instanse):
        if self.text_f1.text==self.ids.tx.text:
            print("hi")
        else:
            print('Bye')
        pass


    def dialog_password(self):

         self.text_f1=MDTextField(
             hint_text="Old password"
         )
         self.text_f2=MDTextField(
             hint_text="New password"
         )
         self.flat_b1=MDFlatButton(
             text="CANCEL",
             theme_text_color="Custom"
         )
         self.flat_b2=MDFlatButton(
             text="OK",
             theme_text_color="Custom",
             on_release=self.change_password
         )
         self.dialog = MDDialog(
                title="Change password:",
                type="custom",
                content_cls=MDBoxLayout(
                    self.text_f1,
                    self.text_f2,
                    orientation="vertical",
                    spacing="12dp",
                    size_hint_y=None,
                    height="120dp",
                ),
                buttons=[
                    self.flat_b1,
                    self.flat_b2
                ],
            )
         self.dialog.open()




class ContentNavigationDrawer(MDBoxLayout):
    pass


class BaseNavigationDrawerItem(MDNavigationDrawerItem):
    def __init__(self, **kwargs):
        super(BaseNavigationDrawerItem,self).__init__(**kwargs)
        self.radius = 24
        self.text_color = "#4a4939"
        self.icon_color = "#4a4939"
        self.focus_color = "#e7e4c0"


class DrawerLabelItem(BaseNavigationDrawerItem):
    def __init__(self, **kwargs):
        super(DrawerLabelItem,self).__init__(**kwargs)
        self.focus_behavior = False
        self._no_ripple_effect = True
        self.selected_color = "#4a4939"


class DrawerClickableItem(BaseNavigationDrawerItem):
    def __init__(self, **kwargs):
        super(DrawerClickableItem,self).__init__(**kwargs)
        self.ripple_color = "#c5bdd2"
        self.selected_color = "#0c6c4d"


class Main2(Screen):

    #hello world
    def __init__(self,**kw):
        super(Main2,self).__init__(**kw)

    def sh(self):
        self.ids.mbutton2.shake()

    def build(self):

        Builder.load_file('ui1.kv')
        pass



class App(MDApp):


    def build(self):

        self.theme_cls.theme_style="Dark"

        self.theme_cls.primary_palette='Red'

        Builder.load_file('ui1.kv')

        sm=ScreenManager()
        sm.add_widget(Main(name="go"))
        sm.add_widget(Main2(name="up"))

        return sm


if __name__=="__main__":
    App().run()