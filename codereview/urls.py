from django.urls import path

from . import views

app_name = 'codereview'

urlpatterns = [
    # index
    path('', views.index, name='index'),
]
