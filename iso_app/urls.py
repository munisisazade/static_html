from django.urls import path
from .views import BaseIndexView


urlpatterns = [
    path('', BaseIndexView.as_view(), name="index"),
]